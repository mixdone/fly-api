package repositories

import (
	"context"
	"fmt"

	"github.com/mixdone/fly-api/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type AirportRepo struct {
	DB  *pgx.Conn
	log *logrus.Logger
}

func NewAirportRepository(db *pgx.Conn, log *logrus.Logger) *AirportRepo {
	return &AirportRepo{
		DB:  db,
		log: log,
	}
}

func (r *AirportRepo) GetAirports() ([]models.Airport, error) {
	rows, err := r.DB.Query(context.Background(), "SELECT airport_code, airport_name, city FROM airports_data")
	if err != nil {
		r.log.Errorf("Error querying airports: %v", err)
		return nil, err
	}
	defer rows.Close()

	var airports []models.Airport
	for rows.Next() {
		var airport models.Airport
		if err := rows.Scan(&airport.Code, &airport.Name, &airport.City); err != nil {
			r.log.Errorf("Error scanning airport: %v", err)
			return nil, err
		}
		airports = append(airports, airport)
	}

	if err := rows.Err(); err != nil {
		r.log.Errorf("Error iterating over rows: %v", err)
		return nil, err
	}

	return airports, nil
}

func (r *AirportRepo) GetAirportsByCity(city string) ([]models.Airport, error) {
	sql := `
		SELECT airport_name, airport_code, city
		FROM airports_data
		WHERE city ->> 'en' = $1;
	`
	r.log.Debug(city)
	rows, err := r.DB.Query(context.Background(), sql, city)

	r.log.Debug(rows)
	if err != nil {
		return nil, fmt.Errorf("failed to query airports: %v", err)
	}
	defer rows.Close()

	var airports []models.Airport

	for rows.Next() {
		var airport models.Airport
		if err := rows.Scan(&airport.Name, &airport.Code, &airport.City); err != nil {
			return nil, fmt.Errorf("failed to scan airport data: %v", err)
		}

		airports = append(airports, airport)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over rows: %v", err)
	}

	return airports, nil
}
