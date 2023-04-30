package fetch

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AdGuardStats struct {
	TimeUnits               string           `json:"time_units"`
	TopQueriedDomains       []map[string]int `json:"top_queried_domains"`
	TopClients              []map[string]int `json:"top_clients"`
	TopBlockedDomains       []map[string]int `json:"top_blocked_domains"`
	DNSQueries              []int            `json:"dns_queries"`
	BlockedFiltering        []int            `json:"blocked_filtering"`
	ReplacedSafeBrowsing    []int            `json:"replaced_safebrowsing"`
	ReplacedParental        []int            `json:"replaced_parental"`
	NumDNSQueries           int              `json:"num_dns_queries"`
	NumBlockedFiltering     int              `json:"num_blocked_filtering"`
	NumReplacedSafeBrowsing int              `json:"num_replaced_safebrowsing"`
	NumReplacedSafeSearch   int              `json:"num_replaced_safesearch"`
	NumReplacedParental     int              `json:"num_replaced_parental"`
	AvgProcessingTime       float64          `json:"avg_processing_time"`
}

func GetAdGuardStats(url, username, password string) (AdGuardStats, error) {
	var stats AdGuardStats

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return stats, fmt.Errorf("failed to create request: %w", err)
	}

	// Add the authorization header
	auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	req.Header.Add("Authorization", "Basic "+auth)

	resp, err := client.Do(req)
	if err != nil {
		return stats, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return stats, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(body, &stats)
	if err != nil {
		return stats, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return stats, nil
}
