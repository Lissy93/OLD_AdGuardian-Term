package pains

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/lissy93/adguardian-term/fetch"
)

func convertToFloatArray(input []int) []float64 {
	floatArr := make([]float64, len(input))
	for i, num := range input {
		floatArr[i] = float64(num)
	}
	return floatArr
}

func AllowedSparkLine(stats fetch.AdGuardStats) *widgets.SparklineGroup {
	sl0 := widgets.NewSparkline()
	sl0.Data = convertToFloatArray(stats.DNSQueries)
	sl0.LineColor = ui.ColorGreen
	slg1 := widgets.NewSparklineGroup(sl0)
	slg1.BorderStyle.Fg = ui.ColorBlue
	slg1.Title = "Allowed Queries"
	return slg1
}

func BlockedSparkLine(stats fetch.AdGuardStats) *widgets.SparklineGroup {
	sl0 := widgets.NewSparkline()
	sl0.Data = convertToFloatArray(stats.BlockedFiltering)
	sl0.LineColor = ui.ColorYellow
	slg1 := widgets.NewSparklineGroup(sl0)
	slg1.BorderStyle.Fg = ui.ColorBlue
	slg1.Title = "Filtered Queries"
	return slg1
}

func ParentalSparkLine(stats fetch.AdGuardStats) *widgets.SparklineGroup {
	sl3 := widgets.NewSparkline()
	sl3.Data = convertToFloatArray(stats.ReplacedParental)
	sl3.LineColor = ui.ColorMagenta
	slg1 := widgets.NewSparklineGroup(sl3)
	slg1.BorderStyle.Fg = ui.ColorBlue
	slg1.Title = "Parental Controls"
	return slg1
}

func MalwareSparkLine(stats fetch.AdGuardStats) *widgets.SparklineGroup {
	sl3 := widgets.NewSparkline()
	sl3.Data = convertToFloatArray(stats.ReplacedSafeBrowsing)
	sl3.LineColor = ui.ColorRed
	slg1 := widgets.NewSparklineGroup(sl3)
	slg1.BorderStyle.Fg = ui.ColorBlue
	slg1.Title = "Malware"
	return slg1
}
