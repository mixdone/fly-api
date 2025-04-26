package repositories

import (
	"github.com/jackc/pgx/v5"
	"github.com/mixdone/fly-api/internal/models"
	"github.com/sirupsen/logrus"
)

type CityRepository interface {
	GetCities() ([]models.City, error)
}

type AirportRepository interface {
	GetAirports() ([]models.Airport, error)
	GetAirportsByCity(city string) ([]models.Airport, error)
}

type Repository struct {
	CityRepository
	AirportRepository
}

func NewRepository(db *pgx.Conn, log *logrus.Logger) *Repository {
	return &Repository{
		CityRepository:    NewCityRepository(db, log),
		AirportRepository: NewAirportRepository(db, log),
	}
}
