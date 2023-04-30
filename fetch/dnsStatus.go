package fetch

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AdGuardStatus struct {
	Version                    string   `json:"version"`
	Language                   string   `json:"language"`
	DNSAddresses               []string `json:"dns_addresses"`
	DNSPort                    int      `json:"dns_port"`
	HTTPPort                   int      `json:"http_port"`
	ProtectionDisabledDuration int      `json:"protection_disabled_duration"`
	ProtectionEnabled          bool     `json:"protection_enabled"`
	DHCPAvailable              bool     `json:"dhcp_available"`
	Running                    bool     `json:"running"`
}

func GetAdGuardStatus(url, username, password string) (AdGuardStatus, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AdGuardStatus{}, err
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encoded))

	resp, err := client.Do(req)
	if err != nil {
		return AdGuardStatus{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AdGuardStatus{}, err
	}

	var status AdGuardStatus
	err = json.Unmarshal(body, &status)
	if err != nil {
		return AdGuardStatus{}, err
	}

	return status, nil
}
