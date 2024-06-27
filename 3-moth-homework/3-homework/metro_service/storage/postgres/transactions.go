package postgres

import (
	"atto/models"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type TransactionRepo struct {
	DB *sql.DB
}

// create
func (c *TransactionRepo) CreateTransaction(transaction *models.CreateTransaction) error {
	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `insert into transactions(id, card_id, amount, terminal_id,
	transaction_type, created_at) 
	values ($1,$2,$3,$4,$5,$6)`
	_, err = tx.Exec(query, uuid.NewString(), transaction.CardId,
		transaction.Amount, transaction.TerminalId,
		transaction.TransactionType, time.Now())
	if err != nil {
		return err
	}
	return nil
}

// read all
func (c *TransactionRepo) GetAllTransactions(filter models.TransactionFilter) (*[]models.Transaction, error) {
	query := `select id, card_id, amount, terminal_id, transaction_type, 
	created_at, updated_at from transactions where and deleted_at is null`
	params := []interface{}{}
	paramcount := 1

	if filter.CardId != nil {
		query += " and card_id = $" + fmt.Sprintf("%d", paramcount)
		params = append(params, *filter.CardId)
		paramcount++
	}
	if filter.Amount != nil {
		query += " and amount = $" + fmt.Sprintf("%d", paramcount)
		params = append(params, *filter.Amount)
		paramcount++
	}
	if filter.TerminalId != nil {
		query += " and terminal_id = $" + fmt.Sprintf("%d", paramcount)
		params = append(params, *filter.TerminalId)
		paramcount++
	}
	if filter.TransactionType != nil {
		query += " and transaction_type = $" + fmt.Sprintf("%d", paramcount)
		params = append(params, *filter.TransactionType)
		paramcount++
	}

	transactions := []models.Transaction{}

	rows, err := c.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		transaction := models.Transaction{}
		err := rows.Scan(&transaction.Id, &transaction.CardId,
			&transaction.Amount, &transaction.TerminalId,
			&transaction.TransactionType, &transaction.CreatedAt,
			&transaction.UpdatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &transactions, nil
}

// read by one
func (c *TransactionRepo) GetTransactionById(transactionId string) (*models.Transaction, error) {
	transaction := models.Transaction{Id: transactionId}

	query := `select id, card_id, amount, terminal_id, transaction_type, 
	created_at, updated_at from transactions where id = $1 and 
	deleted_at is null`

	err := c.DB.QueryRow(query, transactionId).Scan(&transaction.Id,
		&transaction.CardId, &transaction.Amount,
		&transaction.TerminalId, &transaction.TransactionType,
		&transaction.CreatedAt, &transaction.UpdatedAt)

	return &transaction, err
}

// update
func (c *TransactionRepo) UpdateTransaction(transaction *models.UpdateTransaction) error {
	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `update transactions set card_id=$1, amount=$2, terminal_id=$3, 
	transaction_type=$4, updated_at=$5 where id = $6 and deleted_at is null`
	result, err := tx.Exec(query, &transaction.CardId, &transaction.Amount,
		&transaction.TerminalId, &transaction.TransactionType,
		time.Now(), transaction.Id)

	if err != nil {
		return err
	}

	if numberOfAffectedRows, err := result.RowsAffected(); numberOfAffectedRows <= 0 {
		return fmt.Errorf("error: there is no row with the id = %s: %v", transaction.Id, err)
	}
	return nil
}

// delete
func (c *TransactionRepo) DeleteTransaction(transactionID string) error {
	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `update transactions set deleted_at = $1 where id = $2`
	result, err := tx.Exec(query, time.Now(), transactionID)
	if err != nil {
		return err
	}

	if numberOfAffectedRows, err := result.RowsAffected(); numberOfAffectedRows <= 0 {
		return fmt.Errorf("error: there is no row with the id = %s: %s", transactionID, err)
	}
	return nil
}
