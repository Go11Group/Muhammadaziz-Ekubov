package postgres

import (
	"Go11/14-homework/model"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
)

type UserRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func InsertUser(w http.ResponseWriter, r *http.Request) (model.Users, error) {
	if r.Method != http.MethodGet {
		w.Write([]byte("Sorry, incorrect method"))
	}

	tx, err := io.ReadAll(r.Body)
	if err != nil {
		return model.Users{}, err
	}

	var insertUser model.Users
	err = json.Unmarshal(tx, &insertUser)
	if err != nil {
		return model.Users{}, err
	}

	return insertUser, nil
}

func UpdateUser(w http.ResponseWriter, r *http.Request) (model.Users, error) {
	if r.Method != http.MethodGet {
		w.Write([]byte("Sorry, incorrect method"))
	}

	tx, err := io.ReadAll(r.Body)
	if err != nil {
		return model.Users{}, err
	}

	var updateUser model.Users
	err = json.Unmarshal(tx, &updateUser)
	if err != nil {
		return model.Users{}, err
	}

	return updateUser, nil
}

func DeleteUser(w http.ResponseWriter, r *http.Request) (model.Users, error) {
	if r.Method != http.MethodGet {
		w.Write([]byte("Sorry, incorrect method"))
	}

	byt, err := io.ReadAll(r.Body)
	if err != nil {
		return model.Users{}, err
	}

	var deleteUser model.Users
	err = json.Unmarshal(byt, &deleteUser)
	if err != nil {
		return model.Users{}, err
	}

	return deleteUser, nil
}

func InsertProduct(w http.ResponseWriter, r *http.Request) (model.Products, error) {
	if r.Method != http.MethodGet {
		w.Write([]byte("Sorry, incorrect method"))
	}

	tx, err := io.ReadAll(r.Body)
	if err != nil {
		return model.Products{}, err
	}

	var insertProduct model.Products
	err = json.Unmarshal(tx, &insertProduct)
	if err != nil {
		return model.Products{}, err
	}

	return insertProduct, nil
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) (model.Products, error) {
	if r.Method != http.MethodGet {
		w.Write([]byte("Sorry, incorrect method"))
	}

	tx, err := io.ReadAll(r.Body)
	if err != nil {
		return model.Products{}, err
	}

	var productUpdate model.Products
	err = json.Unmarshal(tx, &productUpdate)
	if err != nil {
		return model.Products{}, err
	}

	return productUpdate, nil
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) (model.Products, error) {
	if r.Method != http.MethodGet {
		w.Write([]byte("Bu so'rov get so'rovi emas?"))
	}

	byt, err := io.ReadAll(r.Body)
	if err != nil {
		return model.Products{}, err
	}

	var productDelete model.Products
	err = json.Unmarshal(byt, &productDelete)
	if err != nil {
		return model.Products{}, err
	}

	return productDelete, nil
}
