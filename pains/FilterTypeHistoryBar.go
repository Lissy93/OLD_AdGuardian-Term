package pains

import (
	"fmt"
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/lissy93/adguardian-term/common"
	"github.com/lissy93/adguardian-term/fetch"
)

func generateBarChartData(stats fetch.AdGuardStats) [][]float64 {
	numDays := len(stats.DNSQueries)
	barChartData := make([][]float64, numDays)

	for day := 0; day < numDays; day++ {
		blocked := float64(stats.BlockedFiltering[day])
		safeBrowsing := float64(stats.ReplacedSafeBrowsing[day])
		parental := float64(stats.ReplacedParental[day])
		total := float64(stats.DNSQueries[day]) - blocked - safeBrowsing - parental
		barChartData[day] = []float64{total, blocked, safeBrowsing, parental}
	}

	return barChartData
}

func FilterTypeHistoryBar(stats fetch.AdGuardStats) *widgets.StackedBarChart {
	sbc := widgets.NewStackedBarChart()
	sbc.Title = "Historical Query Categories"
	sbc.Data = generateBarChartData(stats)
	sbc.BarWidth = 4
	sbc.NumFormatter = func(f float64) string {
		return fmt.Sprintf("%.0f", f)
	}

	// Set colors for each category
	sbc.BarColors = []termui.Color{
		termui.ColorGreen,   // Total
		termui.ColorYellow,  // Blocked
		termui.ColorRed,     // Safe Browsing
		termui.ColorMagenta, // Parental
	}
	common.SetCommonStyles(sbc)
	return sbc
}
