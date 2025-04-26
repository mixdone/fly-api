package services

import (
	"github.com/mixdone/fly-api/internal/models"
	"github.com/mixdone/fly-api/internal/repositories"
)

type Airports interface {
	GetAirports() ([]models.Airport, error)
	GetAirportsByCity(city string) ([]models.Airport, error)
}

type Cities interface {
	GetCities() ([]models.City, error)
}

type Service struct {
	Cities
	Airports
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{
		Cities:   NewCityService(repo.CityRepository),
		Airports: NewAirportService(repo.AirportRepository),
	}
}
