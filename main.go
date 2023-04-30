package main

import (
	"log"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/lissy93/adguardian-term/fetch"
	"github.com/lissy93/adguardian-term/pains"
	"github.com/lissy93/adguardian-term/values"
)

type AdGuardStats struct {
	NumDNSQueries           int `json:"num_dns_queries"`
	NumBlockedFiltering     int `json:"num_blocked_filtering"`
	NumReplacedSafeBrowsing int `json:"num_replaced_safebrowsing"`
	NumReplacedParental     int `json:"num_replaced_parental"`
}

func main() {

	const rowCount float64 = 12

	stats, err := fetch.GetAdGuardStats(values.ENDPOINT+"/control/stats", values.USERNAME, values.PASSWORD)
	if err != nil {
		log.Fatalf("failed to fetch AdGuard stats: %v", err)
	}

	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	// Create the first bar chart
	bc1 := widgets.NewBarChart()
	bc1.Data = []float64{10, 20, 30, 40, 50}
	bc1.Title = "Bar Chart 1"
	bc1.Labels = []string{"A", "B", "C", "D", "E"}

	// Set up the grid layout
	termWidth, termHeight := termui.TerminalDimensions()
	grid := termui.NewGrid()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(
		// Row 1 - Title and AdGuard DNS Status
		termui.NewRow(1.0/rowCount,
			termui.NewCol(0.2, pains.Title()),
			termui.NewCol(0.8, pains.DnsStatus()),
		),
		// Row 2 - Block, allow, parental and malware breakdown
		termui.NewRow(3/rowCount,
			termui.NewCol(0.25, pains.BlockPercentage(stats)),
			termui.NewCol(0.25, pains.QueryCount(stats)),
			termui.NewCol(0.5,
				termui.NewCol(0.5,
					termui.NewRow(0.5, pains.AllowedSparkLine(stats)),
					termui.NewRow(0.5, pains.BlockedSparkLine(stats)),
				),
				termui.NewCol(0.5,
					termui.NewRow(0.5, pains.ParentalSparkLine(stats)),
					termui.NewRow(0.5, pains.MalwareSparkLine(stats)),
				),
			),
		),
		// Row 3 - Top queried domains, blocked domains and clients
		termui.NewRow(4.0/rowCount,
			termui.NewCol(0.333, pains.QueryLog()),
			termui.NewCol(0.333, pains.BlockLog()),
			termui.NewCol(0.333, pains.ClientLog()),
		),
		termui.NewRow(3.0/rowCount,
			termui.NewCol(1.0, bc1),
		),
	)

	// Render the grid
	termui.Render(grid)

	// Handle user input
	uiEvents := termui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
