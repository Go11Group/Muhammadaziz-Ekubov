package postgres

import (
	"database/sql"
)

type BusRepo struct {
	DB *sql.DB
}

func NewBusRepo(db *sql.DB) *BusRepo {
	return &BusRepo{
		DB: db,
	}
}

//func (t *BusRepo) GetByBusNumber(busNumber int) (*pb., error) {
//	query := `select * from transports where bus_number = $1`
//
//	Info := pb.BusScheduleRequest{}
//	err := t.DB.QueryRow(query, busNumber).Scan(&info.BusNumber, pq.Array(&info.Stations))
//	return &info, err
//}
