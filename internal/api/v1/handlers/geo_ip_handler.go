package handlers

import (
	"hermes/internal/services/thirdparty"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GeoIPHandler handles IP geolocation requests
type GeoIPHandler struct {
	client *thirdparty.GeoIPClient
}

// NewGeoIPHandler creates a new instance of GeoIPHandler
func NewGeoIPHandler() *GeoIPHandler {
	return &GeoIPHandler{
		client: thirdparty.NewGeoIPClient(),
	}
}

// GetGeoIP retrieves geolocation data for an IP address
func (h *GeoIPHandler) GetGeoIP(c echo.Context) error {
	ip := c.QueryParam("ip")
	if ip == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "IP address is required"})
	}

	geoInfo, err := h.client.GetGeoIP(ip)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, geoInfo)
}
