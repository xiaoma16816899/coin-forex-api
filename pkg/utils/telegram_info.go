package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IPInfoResponse struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
}

func GetTelecomCompany(ip string) (IPInfoResponse, error) {
	resp, err := http.Get("https://ipinfo.io/" + ip + "/json")
	var ipInfo IPInfoResponse
	if err != nil {
		fmt.Printf("Failed to make request: %v\n", err)
		return ipInfo, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&ipInfo)
	if err != nil {
		return ipInfo, err
	}
	// fmt.Printf("Organization (Telecom Company): %s\n", ipInfo)
	return ipInfo, nil
}
