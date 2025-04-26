package services

import (
	"github.com/mixdone/fly-api/internal/models"
	"github.com/mixdone/fly-api/internal/repositories"
)

type CityService struct {
	CityRepo repositories.CityRepository
}

func NewCityService(repo repositories.CityRepository) *CityService {
	return &CityService{
		CityRepo: repo,
	}
}

func (s *CityService) GetCities() ([]models.City, error) {
	return s.CityRepo.GetCities()
}
