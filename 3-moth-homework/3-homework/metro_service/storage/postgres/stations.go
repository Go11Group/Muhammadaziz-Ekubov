package postgres

import (
	"atto/models"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type StationRepo struct {
	DB *sql.DB
}

// create
func (s *StationRepo) CreateStation(station *models.CreateStation) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `insert into stations(id, name, created_at) 
	values ($1,$2,$3)`
	_, err = tx.Exec(query, uuid.NewString(), station.Name,
		time.Now())
	if err != nil {
		return err
	}
	return nil
}

// read all
func (c *StationRepo) GetAllStations(filter models.StationFilter) (*[]models.Station, error) {
	query := `select id, name, created_at, updated_at from 
	stations where and deleted_at is null`
	params := []interface{}{}
	paramcount := 1

	if filter.Name != nil {
		query += " and name = $" + fmt.Sprintf("%d", paramcount)
		params = append(params, *filter.Name)
		paramcount++
	}
	stations := []models.Station{}

	rows, err := c.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		station := models.Station{}
		err := rows.Scan(&station.Id, &station.Name,
			&station.CreatedAt, &station.UpdatedAt)
		if err != nil {
			return nil, err
		}
		stations = append(stations, station)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &stations, nil
}

// read by one
func (c *StationRepo) GetStationById(stationId string) (*models.Station, error) {
	station := models.Station{Id: stationId}

	query := `select name, created_at, updated_at
	from stations where id = $1 and deleted_at is null`

	err := c.DB.QueryRow(query, stationId).Scan(&station.Name,
		&station.CreatedAt, &station.UpdatedAt)
	return &station, err
}

// update
func (s *StationRepo) UpdateStation(station *models.UpdateStation) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `update stations set name=$1, updated_at=$2
	where id = $3 and deleted_at is null`
	result, err := tx.Exec(query, station.Name,
		time.Now(), station.Id)

	if err != nil {
		return err
	}

	if numberOfAffectedRows, err := result.RowsAffected(); numberOfAffectedRows <= 0 {
		return fmt.Errorf("error: there is no row with the id = %s: %v", station.Id, err)
	}
	return nil
}

// delete
func (s *StationRepo) DeleteStation(stationID string) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `update stations set deleted_at = $1 where id = $2`
	result, err := tx.Exec(query, time.Now(), stationID)
	if err != nil {
		return err
	}

	if numberOfAffectedRows, err := result.RowsAffected(); numberOfAffectedRows <= 0 {
		return fmt.Errorf("error: there is no row with the id = %s: %s", stationID, err)
	}
	return nil
}
