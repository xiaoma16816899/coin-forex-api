package openApi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"server.com/app/model"
)

const (
	BASE_URL = "https://api.polygon.io/v2/aggs/ticker"
	API_KEY  = "vgzGx2cEvw2phGWJ8QSanJuS7RIV1MGs"
)

func GetForexExchangeCurrencies(params model.ParamForexTradingModel) (interface{}, error) {
	url := fmt.Sprintf("%s/%s/range/1/day/%s/%s?adjusted=true&sort=asc&limit=%d&apiKey=%s", BASE_URL, params.CurrencyType, params.StartDate, params.EndDate, params.Limit, API_KEY)
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	mapData := map[string]interface{}{}
	if err := json.Unmarshal(body, &mapData); err != nil {
		panic(err)
	}

	data, _ := mapData["results"]
	ticker, _ := mapData["ticker"]

	result := make(map[string]interface{})
	result["result"] = data
	result["ticker"] = ticker
	return result, nil
}

func GetAllForexExchangeCurrencies(date string) (interface{}, error) {
	url := fmt.Sprintf("https://api.polygon.io/v2/aggs/grouped/locale/global/market/fx/%s?adjusted=true&apiKey=%s", date, API_KEY)
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	mapData := map[string]interface{}{}
	if err := json.Unmarshal(body, &mapData); err != nil {
		panic(err)
	}

	data, _ := mapData["results"]

	result := make(map[string]interface{})
	result["result"] = data
	return result, nil
}
