package model

type Users struct {
	Id        int    `json:id`
	FirstName string `json:first_name`
	LastName  string `json:last_name`
	Age       int    `json:age`
}
