package fetch

import (
	"github.com/lissy93/adguardian-term/values"
	"log"
	"time"
)

func FetchData(done chan bool, statsChan chan AdGuardStats, logsChan chan AdGuardQueryLog) {
	for {
		stats, statsErr := GetAdGuardStats(values.ENDPOINT+"/control/stats", values.USERNAME, values.PASSWORD)
		queryLog, queryLogErr := GetAdGuardQueryLog(values.ENDPOINT+"/control/querylog", values.USERNAME, values.PASSWORD)

		if statsErr != nil {
			log.Printf("failed to fetch AdGuard stats: %v", statsErr)
		} else {
			statsChan <- stats
		}

		if queryLogErr != nil {
			log.Printf("failed to fetch AdGuard query log: %v", queryLogErr)
		} else {
			logsChan <- queryLog
		}

		select {
		case <-done:
			return
		case <-time.After(10 * time.Second):
			// Continue fetching data
		}
	}
}
