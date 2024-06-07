package model

type Products struct {
	Id       int     `json:"id" :"id"`
	Name     string  `json:"name" :"name"`
	Price    float64 `json:"price" :"price"`
	Quantity int     `json:"quantity"`
}
