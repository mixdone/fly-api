package repositories

import (
	"context"

	"github.com/mixdone/fly-api/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type CityRepo struct {
	DB  *pgx.Conn
	log *logrus.Logger
}

func NewCityRepository(db *pgx.Conn, log *logrus.Logger) *CityRepo {
	return &CityRepo{
		DB:  db,
		log: log,
	}
}

func (r *CityRepo) GetCities() ([]models.City, error) {
	rows, err := r.DB.Query(context.Background(), "SELECT city FROM airports_data")
	if err != nil {
		r.log.Errorf("Error querying cities: %v", err)
		return nil, err
	}
	defer rows.Close()

	var cities []models.City
	for rows.Next() {
		var city models.City
		if err := rows.Scan(&city.Name); err != nil {
			r.log.Errorf("Error scanning city: %v", err)
			return nil, err
		}
		cities = append(cities, city)
	}

	if err := rows.Err(); err != nil {
		r.log.Errorf("Error iterating over rows: %v", err)
		return nil, err
	}

	return cities, nil
}
