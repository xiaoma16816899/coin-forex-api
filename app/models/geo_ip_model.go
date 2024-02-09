package models

type ClientAddress struct {
	IP          string `json:"ip"`
	Continent   string `json:"continent"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	City        string `json:"city"`
}

type GeoIPRequest struct {
	IP         string `json:"ip"`
	PrivateKey string `json:"privateKey"`
}

type ResponseClientAddress struct {
	Data ClientAddress `json:"data"`
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
}
