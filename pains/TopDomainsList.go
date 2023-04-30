package pains

import (
	"fmt"
	"github.com/gizak/termui/v3/widgets"
	"github.com/lissy93/adguardian-term/common"
	"github.com/lissy93/adguardian-term/fetch"
)

func extractTopQueriedDomainsURLs(domainList []map[string]int, color string) []string {
	urls := make([]string, 0, len(domainList))
	for _, domain := range domainList {
		for url, queryCount := range domain {
			formattedURL := fmt.Sprintf("%s [%d](fg:%s,mod:bold)", url, queryCount, color)
			urls = append(urls, formattedURL)
		}
	}
	return urls
}

func makeLogListWidget(dataList []map[string]int, title string, color string) *widgets.List {
	listWidget := widgets.NewList()
	listWidget.Rows = extractTopQueriedDomainsURLs(dataList, color)
	listWidget.Title = title
	common.SetCommonStyles(listWidget)
	return listWidget
}

func QueryLog(stats fetch.AdGuardStats) *widgets.List {
	return makeLogListWidget(stats.TopQueriedDomains, "Top Queried Domains", "cyan")
}

func BlockLog(stats fetch.AdGuardStats) *widgets.List {
	return makeLogListWidget(stats.TopBlockedDomains, "Top Blocked Domains", "yellow")
}

func ClientLog(stats fetch.AdGuardStats) *widgets.List {
	return makeLogListWidget(stats.TopClients, "Top Clients", "magenta")
}
