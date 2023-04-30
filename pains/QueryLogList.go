package pains

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/lissy93/adguardian-term/fetch"
	"github.com/lissy93/adguardian-term/values"
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

func QueryLog() *widgets.List {
	// Fetch query count stats from AdGuard API
	statsEndpoint := values.ENDPOINT + "/control/stats"
	stats, err := fetch.GetAdGuardStats(statsEndpoint, values.USERNAME, values.PASSWORD)
	if err != nil {
		log.Fatalf("failed to fetch AdGuard stats: %v", err)
	}
	queryCountPain := widgets.NewList()
	queryCountPain.Rows = extractTopQueriedDomainsURLs(stats.TopQueriedDomains, "cyan")
	queryCountPain.Title = "Top Queried Domains"
	queryCountPain.BorderStyle.Fg = ui.ColorBlue

	return queryCountPain
}

func BlockLog() *widgets.List {
	// Fetch query count stats from AdGuard API
	statsEndpoint := values.ENDPOINT + "/control/stats"
	stats, err := fetch.GetAdGuardStats(statsEndpoint, values.USERNAME, values.PASSWORD)
	if err != nil {
		log.Fatalf("failed to fetch AdGuard stats: %v", err)
	}
	queryCountPain := widgets.NewList()
	queryCountPain.Rows = extractTopQueriedDomainsURLs(stats.TopBlockedDomains, "yellow")
	queryCountPain.Title = "Top Blocked Domains"
	queryCountPain.BorderStyle.Fg = ui.ColorBlue

	return queryCountPain
}

func ClientLog() *widgets.List {
	// Fetch query count stats from AdGuard API
	statsEndpoint := values.ENDPOINT + "/control/stats"
	stats, err := fetch.GetAdGuardStats(statsEndpoint, values.USERNAME, values.PASSWORD)
	if err != nil {
		log.Fatalf("failed to fetch AdGuard stats: %v", err)
	}
	queryCountPain := widgets.NewList()
	queryCountPain.Rows = extractTopQueriedDomainsURLs(stats.TopClients, "magenta")
	queryCountPain.Title = "Top Clients"
	queryCountPain.BorderStyle.Fg = ui.ColorBlue

	return queryCountPain
}
