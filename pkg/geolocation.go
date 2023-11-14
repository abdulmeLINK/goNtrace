package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GeoLocation struct {
	Country string  `json:"country"`
	City    string  `json:"city"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lon"`
	Query   string  `json:"query"`
}

func GetGeoLocation(ip string) (GeoLocation, error) {
	resp, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s", ip))
	if err != nil {
		return GeoLocation{}, fmt.Errorf("could not get geolocation: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GeoLocation{}, fmt.Errorf("could not read response body: %v", err)
	}

	var geoLocation GeoLocation
	err = json.Unmarshal(body, &geoLocation)
	if err != nil {
		return GeoLocation{}, fmt.Errorf("could not unmarshal JSON: %v", err)
	}

	return geoLocation, nil
}
