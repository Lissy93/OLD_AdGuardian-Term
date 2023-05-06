package pains

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/lissy93/adguardian-term/common"
	"github.com/lissy93/adguardian-term/fetch"
	"math"
	"time"
)

func queriesPerMinute(logs fetch.AdGuardQueryLog) (map[int64]int, map[int64]int) {
	blockedCounts := make(map[int64]int)
	allowedCounts := make(map[int64]int)
	now := time.Now()

	for _, log := range logs.Data {
		logTime, err := time.Parse(time.RFC3339, log.Time)
		if err != nil {
			continue
		}
		minutesAgo := int64(now.Sub(logTime).Minutes())
		if log.Reason != "NotFilteredNotFound" {
			blockedCounts[minutesAgo]++
		} else {
			allowedCounts[minutesAgo]++
		}
	}

	return allowedCounts, blockedCounts
}

func prepareScatterPlotData(logs fetch.AdGuardQueryLog, plotWidth int) [][]float64 {
	allowedRequestCount, blockedRequestCounts := queriesPerMinute(logs)

	stretchedData := make([][]float64, 2)
	stretchedData[0] = stretchData(allowedRequestCount, plotWidth)
	stretchedData[1] = stretchData(blockedRequestCounts, plotWidth)

	return stretchedData
}

func stretchData(data map[int64]int, targetLength int) []float64 {
	stretched := []float64{}

	originalLength := len(data)
	if originalLength == 0 {
		return stretched
	}

	stretchFactor := float64(targetLength) / float64(originalLength)

	for _, count := range data {
		repeatCount := int(math.Round(stretchFactor))
		for i := 0; i < repeatCount; i++ {
			stretched = append(stretched, float64(count))
		}
	}

	return stretched
}

func QueryTimeLine(queryLog fetch.AdGuardQueryLog, width int) *widgets.Plot {
	p0 := widgets.NewPlot()
	p0.Title = fmt.Sprintf("Current Query Volume %d", width)
	p0.Data = make([][]float64, 0)
	p0.Data = prepareScatterPlotData(queryLog, width)
	p0.LineColors[0] = ui.ColorGreen
	p0.LineColors[1] = ui.ColorRed
	common.SetCommonStyles(p0)
	p0.DataLabels = []string{"Allowed", "Blocked"}
	p0.DrawDirection = widgets.DrawRight
	p0.MaxVal = 50
	return p0
}
