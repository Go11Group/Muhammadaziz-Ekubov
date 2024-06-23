package CRUDS

import (
	"Go11/14-homework/model"
	"Go11/14-homework/storage/postgres"
	"database/sql"
	"encoding/json"
	"net/http"
)

type CrudProduct struct {
	db *sql.DB
}

func NewCrudProductRepo(db *sql.DB) *CrudProduct {
	return &CrudProduct{db: db}
}

func (i *CrudProduct) InsertProduct(w http.ResponseWriter, r *http.Request) {
	product, err := postgres.InsertProduct(w, r)
	if err != nil {
		panic(err)
	}

	_, err = i.db.Exec("insert into products(id,name,price,quantity) values($1,$2,$3,$4)",
		&product.Id, &product.Name, &product.Price, &product.Quantity)
	if err != nil {
		panic(err)
	}
}

func (u *CrudProduct) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	up, err := postgres.UpdateProduct(w, r)
	if err != nil {
		panic(err)
	}

	_, err = u.db.Exec("update products set name=$1, price=$2,quantity=&3  where id=$4",
		&up.Name, &up.Price, &up.Quantity, &up.Id)

	if err != nil {
		panic(err)
	}
}

func (d *CrudProduct) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	del, err := postgres.DeleteProduct(w, r)
	if err != nil {
		panic(err)
	}

	_, err = d.db.Exec("delete from products where id=$1", &del.Id)
	if err != nil {
		panic(err)
	}
}

func (s *CrudProduct) ReadProduct(w http.ResponseWriter, r *http.Request) {
	row, err := s.db.Query("select * from products")
	if err != nil {
		http.Error(w, "Unable to scan the row", http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var products []model.Products
	for row.Next() {
		var product model.Products
		err = row.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity)
		if err != nil {
			http.Error(w, "Unable to scan the row", http.StatusInternalServerError)
			return
		}

		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		panic(err)
	}
}
