package thirdparty

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GeoIPInfo represents the geolocation data returned from ip-api.com
type GeoIPInfo struct {
	IP          string  `json:"query"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	AS          string  `json:"as"`
}

// GeoIPClient handles API requests to ip-api.com
type GeoIPClient struct {
	BaseURL string
	Client  *http.Client
}

// NewGeoIPClient creates a new instance of the geolocation client
func NewGeoIPClient() *GeoIPClient {
	return &GeoIPClient{
		BaseURL: "http://ip-api.com/json/",
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// GetGeoIP fetches geolocation data for a given IP address
func (g *GeoIPClient) GetGeoIP(ip string) (*GeoIPInfo, error) {
	url := fmt.Sprintf("%s%s", g.BaseURL, ip)

	resp, err := g.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch geolocation data")
	}

	var geoData GeoIPInfo
	if err := json.NewDecoder(resp.Body).Decode(&geoData); err != nil {
		return nil, err
	}

	return &geoData, nil
}
