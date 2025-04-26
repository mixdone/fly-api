package services

import (
	"github.com/mixdone/fly-api/internal/models"
	"github.com/mixdone/fly-api/internal/repositories"
)

type AirportService struct {
	AirportRepo repositories.AirportRepository
}

func NewAirportService(repo repositories.AirportRepository) *AirportService {
	return &AirportService{
		AirportRepo: repo,
	}
}

func (s *AirportService) GetAirports() ([]models.Airport, error) {
	return s.AirportRepo.GetAirports()
}

func (s *AirportService) GetAirportsByCity(city string) ([]models.Airport, error) {
	return s.AirportRepo.GetAirportsByCity(city)
}
