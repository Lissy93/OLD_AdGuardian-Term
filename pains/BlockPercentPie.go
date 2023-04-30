package pains

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/lissy93/adguardian-term/fetch"
)

func BlockPercentage(stats fetch.AdGuardStats) *widgets.PieChart {
	// Create the pie chart
	chartBlockPercent := widgets.NewPieChart()
	chartBlockPercent.Title = "Block Percentage"
	chartBlockPercent.Data = []float64{
		float64(stats.NumBlockedFiltering),
		float64(stats.NumDNSQueries - stats.NumBlockedFiltering - stats.NumReplacedParental - stats.NumReplacedSafeBrowsing),
		float64(stats.NumReplacedSafeBrowsing),
		float64(stats.NumReplacedParental),
	}
	total := float64(stats.NumDNSQueries)
	chartBlockPercent.LabelFormatter = func(i int, v float64) string {
		percentage := (v / total) * 100
		var label string
		switch i {
		case 0:
			label = "Blocked"
		case 1:
			label = "Allowed"
		case 2:
			label = "Malware"
		default:
			label = "Other"
		}
		return fmt.Sprintf("%s %.02f%%", label, percentage)
	}
	chartBlockPercent.AngleOffset = 90
	chartBlockPercent.BorderStyle.Fg = ui.ColorBlue
	return chartBlockPercent
}

func QueryCount(stats fetch.AdGuardStats) *widgets.Table {
	queryCountPain := widgets.NewTable()
	queryCountPain.Rows = [][]string{
		{"Allowed", fmt.Sprintf("[%d](fg:green,mod:bold)", stats.NumDNSQueries-stats.NumBlockedFiltering-stats.NumReplacedParental-stats.NumReplacedSafeBrowsing)},
		{"Filtered", fmt.Sprintf("[%d](fg:yellow,mod:bold)", stats.NumBlockedFiltering)},
		{"Parental Controls", fmt.Sprintf("[%d](fg:magenta,mod:bold)", stats.NumReplacedParental)},
		{"Malware", fmt.Sprintf("[%d](fg:red,mod:bold)", stats.NumReplacedSafeBrowsing)},
		{"Total", fmt.Sprintf("[%d](fg:cyan,mod:bold)", stats.NumDNSQueries)},
	}
	queryCountPain.TextStyle = ui.NewStyle(ui.ColorWhite)
	queryCountPain.Title = "Query Count"
	queryCountPain.TextAlignment = ui.AlignCenter
	queryCountPain.BorderStyle.Fg = ui.ColorBlue
	queryCountPain.FillRow = true

	return queryCountPain
}
