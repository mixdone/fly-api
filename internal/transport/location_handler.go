package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/mixdone/fly-api/internal/services"
)

type LocationsHandler struct {
	CityService    services.Cities
	AirportService services.Airports
}

// GetCities godoc
// @Summary Получить список городов
// @Tags Locations
// @Produce json
// @Success 200 {array} models.City
// @Failure 500
// @Router /locations/cities [get]
func (h *LocationsHandler) GetCities(c *gin.Context) {
	cities, err := h.CityService.GetCities()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, cities)
}

// GetAirports godoc
// @Summary Получить список аэропортов
// @Tags Locations
// @Produce json
// @Success 200 {array} models.Airport
// @Failure 500
// @Router /locations/airports [get]
func (h *LocationsHandler) GetAirports(c *gin.Context) {
	airports, err := h.AirportService.GetAirports()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, airports)
}

// GetAirportsByCity godoc
// @Summary Получить аэропорты по городу
// @Description Этот эндпоинт позволяет получить список аэропортов по имени города города(en).
// @Tags Locations
// @Produce json
// @Param city path string true "Имя города на английском	"
// @Success 200 {array} models.Airport
// @Failure 500 {object} map[string]string
// @Router /locations/airports/city/{city} [get]
func (h *LocationsHandler) GetAirportsByCity(c *gin.Context) {
	city := c.Param("city")

	airports, err := h.AirportService.GetAirportsByCity(city)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, airports)
}
