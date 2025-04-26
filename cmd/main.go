package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mixdone/fly-api/docs"
	"github.com/mixdone/fly-api/internal/config"
	"github.com/mixdone/fly-api/internal/database"
	"github.com/mixdone/fly-api/internal/repositories"
	"github.com/mixdone/fly-api/internal/services"
	"github.com/mixdone/fly-api/internal/transport"
	"github.com/mixdone/fly-api/pkg/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Fly API
// @version 1.0
// @description Это FLY API
// @host localhost:8080
// @BasePath /
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	log := logger.NewLogger(cfg)

	db, err := database.ConnectToDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close(context.Background())

	repository := repositories.NewRepository(db, log)
	service := services.NewService(repository)

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	locationsHandler := &transport.LocationsHandler{
		CityService:    service.Cities,
		AirportService: service.Airports,
	}

	r.GET("/locations/cities", locationsHandler.GetCities)
	r.GET("/locations/airports", locationsHandler.GetAirports)
	r.GET("/locations/airports/city/:city", locationsHandler.GetAirportsByCity)

	r.Run(":8080")

}
