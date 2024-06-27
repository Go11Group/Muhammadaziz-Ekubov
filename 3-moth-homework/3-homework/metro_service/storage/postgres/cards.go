package postgres

import (
	"atto/models"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CardRepo struct {
	DB *sql.DB
}

// create
func (c *CardRepo) CreateCard(card *models.CreateCard) error {
	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `insert into cards(id, number, user_id, created_at) 
	values ($1,$2,$3,$4)`
	_, err = tx.Exec(query, uuid.NewString(), card.Number, card.UserId,
		time.Now())
	if err != nil {
		return err
	}
	return nil
}

// read all
func (c *CardRepo) GetAllCards(filter models.CardFilter) (*[]models.Card, error) {
	query := `select id, number, user_id, created_at, updated_at from 
	cards where and deleted_at is null`
	params := []interface{}{}
	paramcount := 1

	if filter.Number != nil {
		query += " and number = $" + fmt.Sprintf("%d", paramcount)
		params = append(params, *filter.Number)
		paramcount++
	}
	if filter.UserId != nil {
		query += " and user_id = $" + fmt.Sprintf("%d", paramcount)
		params = append(params, *filter.UserId)
		paramcount++
	}
	cards := []models.Card{}

	rows, err := c.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		card := models.Card{}
		err := rows.Scan(&card.Id, &card.Number, &card.UserId,
			&card.CreatedAt, &card.UpdatedAt)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &cards, nil
}

// read by one
func (c *CardRepo) GetCardById(cardId string) (*models.Card, error) {
	card := models.Card{Id: cardId}

	query := `select number, user_id, created_at, updated_at
	from cards where id = $1 and deleted_at is null`

	err := c.DB.QueryRow(query, cardId).Scan(&card.Number, &card.UserId,
		&card.CreatedAt, &card.UpdatedAt)
	return &card, err
}

// update
func (c *CardRepo) UpdateCard(card *models.UpdateCard) error {
	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `update cards set number=$1, user_id=$2, updated_at=$4
	where id = $5 and deleted_at is null`
	result, err := tx.Exec(query, card.Number, card.UserId,
		time.Now(), card.Id)

	if err != nil {
		return err
	}

	if numberOfAffectedRows, err := result.RowsAffected(); numberOfAffectedRows <= 0 {
		return fmt.Errorf("error: there is no row with the id = %s: %v", card.Id, err)
	}
	return nil
}

// delete
func (c *CardRepo) DeleteCard(cardID string) error {
	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `update cards set deleted_at = $1 where id = $2`
	result, err := tx.Exec(query, time.Now(), cardID)
	if err != nil {
		return err
	}

	if numberOfAffectedRows, err := result.RowsAffected(); numberOfAffectedRows <= 0 {
		return fmt.Errorf("error: there is no row with the id = %s: %s", cardID, err)
	}
	return nil
}
