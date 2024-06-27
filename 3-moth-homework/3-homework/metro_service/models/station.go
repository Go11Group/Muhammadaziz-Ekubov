package models

type Station struct {
	Id   string `json:"id"`
	Name int    `json:"name"`
	Time
}

type CreateStation struct {
	Name int `json:"name"`
}

type UpdateStation struct {
	Id   string `json:"id"`
	Name int    `json:"name"`
}

type StationFilter struct {
	Name *int
}
