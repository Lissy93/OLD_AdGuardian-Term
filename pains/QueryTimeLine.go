package pains

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/lissy93/adguardian-term/common"
	"github.com/lissy93/adguardian-term/fetch"
	"time"
)

func queriesPerMinute(logs fetch.AdGuardQueryLog) (map[int64]int, map[int64]int) {
	counts := make(map[int64]int)
	blockedCounts := make(map[int64]int)
	allowedCounts := make(map[int64]int)
	now := time.Now()

	for _, log := range logs.Data {
		logTime, err := time.Parse(time.RFC3339, log.Time)
		if err != nil {
			continue
		}
		minutesAgo := int64(now.Sub(logTime).Minutes())
		if log.Reason == "FilteredBlackList" {
			blockedCounts[minutesAgo]++
			allowedCounts[minutesAgo] += 0
		} else {
			allowedCounts[minutesAgo]++
			blockedCounts[minutesAgo] += 0
		}
		counts[minutesAgo]++
	}

	return allowedCounts, blockedCounts
}

func prepareScatterPlotData(logs fetch.AdGuardQueryLog) [][]float64 {
	data := make([][]float64, 2)
	data[0] = []float64{}
	data[1] = []float64{}

	allowedRequestCount, blockedRequestCounts := queriesPerMinute(logs)

	for _, count := range allowedRequestCount {
		data[0] = append(data[0], float64(count))
	}
	for _, count := range blockedRequestCounts {
		data[1] = append(data[1], float64(count))
	}

	return data
}

func QueryTimeLine(queryLog fetch.AdGuardQueryLog) *widgets.Plot {
	p0 := widgets.NewPlot()
	p0.Title = "Current Query Volume"
	p0.Data = prepareScatterPlotData(queryLog)
	p0.LineColors[0] = ui.ColorGreen
	p0.LineColors[1] = ui.ColorRed
	common.SetCommonStyles(p0)
	p0.DataLabels = []string{"Allowed", "Blocked"}
	p0.DrawDirection = widgets.DrawRight
	return p0
}
