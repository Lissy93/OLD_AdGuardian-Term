package main

import (
	"github.com/gizak/termui/v3"
	ui "github.com/gizak/termui/v3"
	"github.com/lissy93/adguardian-term/fetch"
	"github.com/lissy93/adguardian-term/pains"
	"github.com/lissy93/adguardian-term/values"
	"log"
	"time"
)

func main() {

	const rowCount float64 = 12

	stats, statsErr := fetch.GetAdGuardStats(values.ENDPOINT+"/control/stats", values.USERNAME, values.PASSWORD)
	if statsErr != nil {
		log.Fatalf("failed to fetch AdGuard stats: %v", statsErr)
	}

	queryLog, queryLogErr := fetch.GetAdGuardQueryLog(values.ENDPOINT+"/control/querylog", values.USERNAME, values.PASSWORD)
	if queryLogErr != nil {
		log.Fatalf("failed to fetch AdGuard query log: %v", queryLogErr)
	}

	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	// Widgets to render
	titleWidget := pains.Title()
	statusWidget := pains.DnsStatus()
	pieChartWidget := pains.BlockPercentage(stats)
	queryCountWidget := pains.QueryCount(stats)
	allowSparkWidget := pains.AllowedSparkLine(stats)
	blockedSparkWidget := pains.BlockedSparkLine(stats)
	parentalSparkWidget := pains.ParentalSparkLine(stats)
	malwareSparkWidget := pains.MalwareSparkLine(stats)
	queryLogWidget := pains.QueryLog(stats)
	blockLogWidget := pains.BlockLog(stats)
	clientLogWidget := pains.ClientLog(stats)
	queryTreeWidget := pains.QueryTree(queryLog)
	queryTimePlot := pains.QueryTimeLine(queryLog)

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

	// Set up the grid layout
	termWidth, termHeight := termui.TerminalDimensions()
	grid := termui.NewGrid()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(row1, row2, row3, row4)

	// Render the grid
	termui.Render(grid)

	// Timer, for re-fetching data + updating UI
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
				ui.Clear()
			case "r":
				ui.Clear()
				ui.Render(grid)
			case "<Enter>":
				queryTreeWidget.ToggleExpand()
				ui.Render(queryTreeWidget)
			case "j", "<Down>":
				queryTreeWidget.ScrollDown()
				ui.Render(queryTreeWidget)
			case "k", "<Up>":
				queryTreeWidget.ScrollUp()
				ui.Render(queryTreeWidget)
			}
		}
	}

}
