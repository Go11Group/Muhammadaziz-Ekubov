package CRUD

import (
	"database/sql"
)

// Transaction

func CreateUserAndProduct(db *sql.DB, username, email, password, productName, productDescription string, price float64, stockQuantity int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	var userID, productID int

	err = tx.QueryRow("INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id", username, email, password).Scan(&userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.QueryRow("INSERT INTO products (name, description, price, stock_quantity) VALUES ($1, $2, $3, $4) RETURNING id", productName, productDescription, price, stockQuantity).Scan(&productID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("INSERT INTO user_products (user_id, product_id) VALUES ($1, $2)", userID, productID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func UpdateUserAndProduct(db *sql.DB, userID int, username, email, password string, productID int, productName, productDescription string, price float64, stockQuantity int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4", username, email, password, userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE products SET name = $1, description = $2, price = $3, stock_quantity = $4 WHERE id = $5", productName, productDescription, price, stockQuantity, productID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func DeleteUserAndProduct(db *sql.DB, userID, productID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM user_products WHERE user_id = $1 OR product_id = $2", userID, productID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM products WHERE id = $1", productID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
