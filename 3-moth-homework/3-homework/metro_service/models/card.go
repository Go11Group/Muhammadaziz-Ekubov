package models

type Card struct {
	Id     string `json:"id"`
	Number int    `json:"number"`
	UserId string `json:"user_id"`
	Time
}

type CreateCard struct {
	Number int    `json:"number"`
	UserId string `json:"user_id"`
}

type UpdateCard struct {
	Id     string `json:"id"`
	Number int    `json:"number"`
	UserId string `json:"user_id"`
}

type CardFilter struct {
	Number *int
	UserId *string
}
