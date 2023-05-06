package main

import (
	"github.com/gizak/termui/v3"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/lissy93/adguardian-term/fetch"
	"github.com/lissy93/adguardian-term/pains"
	"github.com/lissy93/adguardian-term/values"
	tb "github.com/nsf/termbox-go"
	"log"
	"math"
	"time"
)

func main() {

	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	tb.SetInputMode(tb.InputEsc)

	// Create the channels
	done := make(chan bool)
	statsChan := make(chan fetch.AdGuardStats)
	logsChan := make(chan fetch.AdGuardQueryLog)

	// Start the fetchData function as a goroutine
	go fetch.FetchData(done, statsChan, logsChan)

	const rowCount float64 = 12

	// Define widgets
	var (
		grid                *termui.Grid
		titleWidget         *widgets.Paragraph
		statusWidget        *widgets.Paragraph
		pieChartWidget      *widgets.PieChart
		queryCountWidget    *widgets.Table
		allowSparkWidget    *widgets.SparklineGroup
		blockedSparkWidget  *widgets.SparklineGroup
		parentalSparkWidget *widgets.SparklineGroup
		malwareSparkWidget  *widgets.SparklineGroup
		queryLogWidget      *widgets.List
		blockLogWidget      *widgets.List
		clientLogWidget     *widgets.List
		queryTreeWidget     *widgets.Tree
		queryTimePlot       *widgets.Plot
		termWidth           int
		termHeight          int
	)

	setWidgetData := func(stats fetch.AdGuardStats, queryLog fetch.AdGuardQueryLog) {
		titleWidget = pains.Title()
		statusWidget = pains.DnsStatus()
		pieChartWidget = pains.BlockPercentage(stats)
		if stats.NumDNSQueries > 0 {
			queryCountWidget = pains.QueryCount(stats)
		}
		allowSparkWidget = pains.AllowedSparkLine(stats)
		blockedSparkWidget = pains.BlockedSparkLine(stats)
		parentalSparkWidget = pains.ParentalSparkLine(stats)
		malwareSparkWidget = pains.MalwareSparkLine(stats)
		queryLogWidget = pains.QueryLog(stats)
		blockLogWidget = pains.BlockLog(stats)
		clientLogWidget = pains.ClientLog(stats)
		queryTreeWidget = pains.QueryTree(queryLog)
		if len(queryLog.Data) > 0 {
			const queryPlotProportion float64 = 0.6 * 3 // Because takes 0.6 of screen, and each plot is 3 chars wide
			queryTimePlot = pains.QueryTimeLine(queryLog, int(math.Round(float64(termWidth)*queryPlotProportion)))
		}
	}

	termWidth, termHeight = termui.TerminalDimensions()
	updateWidgetData := func() {
		stats, statsErr := fetch.GetAdGuardStats(values.ENDPOINT+"/control/stats", values.USERNAME, values.PASSWORD)
		if statsErr != nil {
			log.Printf("failed to fetch AdGuard stats: %v", statsErr)
		}

		queryLog, queryLogErr := fetch.GetAdGuardQueryLog(values.ENDPOINT+"/control/querylog", values.USERNAME, values.PASSWORD)
		if queryLogErr != nil {
			log.Printf("failed to fetch AdGuard query log: %v", queryLogErr)
		}

		setWidgetData(stats, queryLog)
	}

	termWidth, termHeight = termui.TerminalDimensions()
	updateWidgetData()

	grid = termui.NewGrid()

	// Sets up the grid, with widgets
	renderWidgets := func() {
		// Set up the grid layout
		termWidth, termHeight = termui.TerminalDimensions()
		grid.SetRect(0, 0, termWidth, termHeight)

		// Row 1 - Title and AdGuard DNS Status
		row1 := termui.NewRow(
			1.0/rowCount,
			termui.NewCol(0.2, titleWidget),
			termui.NewCol(0.8, statusWidget),
		)

		// Row 2 - Block, allow, parental and malware breakdown
		row2 := termui.NewRow(3/rowCount,
			termui.NewCol(0.25, pieChartWidget),
			termui.NewCol(0.25, queryCountWidget),
			termui.NewCol(0.5,
				termui.NewCol(0.5,
					termui.NewRow(0.5, allowSparkWidget),
					termui.NewRow(0.5, blockedSparkWidget),
				),
				termui.NewCol(0.5,
					termui.NewRow(0.5, parentalSparkWidget),
					termui.NewRow(0.5, malwareSparkWidget),
				),
			),
		)

		// Row 3 - Top queried domains, blocked domains and clients
		row3 := termui.NewRow(4.0/rowCount,
			termui.NewCol(0.333, queryLogWidget),
			termui.NewCol(0.333, blockLogWidget),
			termui.NewCol(0.333, clientLogWidget),
		)

		// Row 4 - Query tree
		row4 := termui.NewRow(4.0/rowCount,
			termui.NewCol(0.25, queryTreeWidget),
			termui.NewCol(0.75, queryTimePlot),
		)

		grid.Set(row1, nil, row2, row3, row4)

		// Render the grid
		termui.Render(grid)
	}

	renderWidgets()

	uiEvents := ui.PollEvents()
	var stats fetch.AdGuardStats
	var queryLog fetch.AdGuardQueryLog

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case stats := <-statsChan:
			setWidgetData(stats, queryLog)
		case logs := <-logsChan:
			setWidgetData(stats, logs)
		case <-ticker.C:
			renderWidgets()
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				renderWidgets()
			case "r":
				renderWidgets()
			case "<Enter>":
				queryTreeWidget.ToggleExpand()
				termui.Render(queryTreeWidget)
			case "j", "<Down>":
				queryTreeWidget.ScrollDown()
				termui.Render(queryTreeWidget)
			case "k", "<Up>":
				queryTreeWidget.ScrollUp()
				termui.Render(queryTreeWidget)
			}
		}
	}
}
