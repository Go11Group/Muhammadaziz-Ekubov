package handler

import (
	crud "Go11/14-homework/CRUDS"
	"net/http"
)

type HandlerUser struct {
	User    *crud.CrudUser
	Product *crud.CrudProduct
}

type HandlerProduct struct {
	Product *crud.CrudProduct
}

func NewHandler(handler HandlerUser) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/insertuser", handler.User.InsertUsers)
	mux.HandleFunc("/updateuser", handler.User.UpdateUsers)
	mux.HandleFunc("/deleteuser", handler.User.DeleteUsers)
	mux.HandleFunc("/readuser", handler.User.ReadUsers)
	mux.HandleFunc("/insertproduct", handler.Product.InsertProduct)
	mux.HandleFunc("/updateproduct", handler.Product.UpdateProduct)
	mux.HandleFunc("/dalateproduct", handler.Product.DeleteProduct)
	mux.HandleFunc("/readproduct", handler.Product.ReadProduct)

	return &http.Server{Handler: mux}
}
