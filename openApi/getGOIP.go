package openAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	models "server.com/app/models"
)

func RouteGEO_IP_API(endPoint string, body models.GeoIPRequest) (models.ResponseClientAddress, error) {
	url := fmt.Sprintf(os.Getenv("APP_GEOIPDOMAIN") + endPoint)
	var data models.ResponseClientAddress
	geoIP := new(models.GeoIPRequest)
	geoIP.IP = body.IP
	geoIP.PrivateKey = body.PrivateKey
	jsonGeoIP, _ := json.Marshal(geoIP)

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonGeoIP))
	if err != nil {
		return data, err
	}

	defer res.Body.Close()
	b, _ := io.ReadAll(res.Body)

	json.Unmarshal(b, &data)
	return data, nil
}