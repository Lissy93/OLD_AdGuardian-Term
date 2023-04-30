package fetch

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AdGuardQueryLog struct {
	Data []struct {
		Answer []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
			TTL   int    `json:"ttl"`
		} `json:"answer"`
		AnswerDNSSEC bool   `json:"answer_dnssec"`
		Cached       bool   `json:"cached"`
		Client       string `json:"client"`
		ClientInfo   struct {
			Whois          json.RawMessage `json:"whois"`
			Name           string          `json:"name"`
			DisallowedRule string          `json:"disallowed_rule"`
			Disallowed     bool            `json:"disallowed"`
		} `json:"client_info"`
		ClientProto string `json:"client_proto"`
		ElapsedMs   string `json:"elapsedMs"`
		Question    struct {
			Class string `json:"class"`
			Name  string `json:"name"`
			Type  string `json:"type"`
		} `json:"question"`
		Reason   string        `json:"reason"`
		Rules    []interface{} `json:"rules"`
		Status   string        `json:"status"`
		Time     string        `json:"time"`
		Upstream string        `json:"upstream"`
	} `json:"data"`
}

func GetAdGuardQueryLog(url, username, password string) (AdGuardQueryLog, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AdGuardQueryLog{}, fmt.Errorf("failed to create request: %w", err)
	}

	encodedAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	req.Header.Add("Authorization", "Basic "+encodedAuth)

	resp, err := client.Do(req)
	if err != nil {
		return AdGuardQueryLog{}, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return AdGuardQueryLog{}, fmt.Errorf("failed to fetch data, status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AdGuardQueryLog{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var queryLog AdGuardQueryLog
	if err := json.Unmarshal(body, &queryLog); err != nil {
		return AdGuardQueryLog{}, fmt.Errorf("failed to unmarshal JSON data: %w", err)
	}

	return queryLog, nil
}
