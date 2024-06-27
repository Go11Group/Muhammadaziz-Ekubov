package postgres

import (
	"atto/models"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type TerminalRepo struct {
	DB *sql.DB
}

// create
func (t *TerminalRepo) CreateTerminal(terminal *models.CreateTerminal) error {
	tx, err := t.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `insert into terminals(id, station_id, created_at) 
	values ($1,$2,$3)`
	_, err = tx.Exec(query, uuid.NewString(), terminal.StationId,
		time.Now())
	if err != nil {
		return err
	}
	return nil
}

// read all
func (t *TerminalRepo) GetAllTerminals(filter models.TerminalFilter) (*[]models.Terminal, error) {
	query := `select id, station_id, created_at, updated_at from 
	terminals where and deleted_at is null`
	params := []interface{}{}
	paramcount := 1

	if filter.StationId != nil {
		query += " and station_id = $" + fmt.Sprintf("%d", paramcount)
		params = append(params, *filter.StationId)
		paramcount++
	}

	terminals := []models.Terminal{}

	rows, err := t.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		terminal := models.Terminal{}
		err := rows.Scan(&terminal.Id, &terminal.StationId,
			&terminal.CreatedAt, &terminal.UpdatedAt)
		if err != nil {
			return nil, err
		}
		terminals = append(terminals, terminal)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &terminals, nil
}

// read by one
func (t *TerminalRepo) GetTerminalById(terminalId string) (*models.Terminal, error) {
	terminal := models.Terminal{Id: terminalId}

	query := `select station_id, created_at, updated_at
	from terminals where id = $1 and deleted_at is null`

	err := t.DB.QueryRow(query, terminalId).Scan(&terminal.StationId,
		&terminal.CreatedAt, &terminal.UpdatedAt)
	return &terminal, err
}

// update
func (t *TerminalRepo) UpdateTerminal(terminal *models.UpdateTerminal) error {
	tx, err := t.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `update terminals set station_id=$1, updated_at=$2
	where id = $3 and deleted_at is null`
	result, err := tx.Exec(query, terminal.StationId,
		time.Now(), terminal.Id)

	if err != nil {
		return err
	}

	if numberOfAffectedRows, err := result.RowsAffected(); numberOfAffectedRows <= 0 {
		return fmt.Errorf("error: there is no row with the id = %s: %v", terminal.Id, err)
	}
	return nil
}

// delete
func (t *TerminalRepo) DeleteTerminal(terminalID string) error {
	tx, err := t.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `update terminals set deleted_at = $1 where id = $2`
	result, err := tx.Exec(query, time.Now(), terminalID)
	if err != nil {
		return err
	}

	if numberOfAffectedRows, err := result.RowsAffected(); numberOfAffectedRows <= 0 {
		return fmt.Errorf("error: there is no row with the id = %s: %s", terminalID, err)
	}
	return nil
}
