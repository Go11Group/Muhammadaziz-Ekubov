package posgres

import (
	pbt "Muhammadaziz-Ekubov/3-moth-homework/7-homework/api-getaway/genproto/weather_service"
	"database/sql"
	time2 "time"
)

type Weather struct {
	DB *sql.DB
}

func NewWeather(db *sql.DB) *Weather {
	return &Weather{
		DB: db,
	}
}

func (w Weather) GetCurrentWeather(city string) (*pbt.GCWResponse, error) {
	res := pbt.GCWResponse{}
	query := `select * from Weather where City = $1`

	var time time2.Time
	time = time2.Now()

	row := w.DB.QueryRow(query, city)
	err := row.Scan(&time, &res.Humidity, &res.Temperature, &res.WindMPH)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (w Weather) CreateWeatherForecast(req *pbt.GWFRequest) (bool, error) {
	query := `insert into Weather (time , humiddity, temperature, windmph) values ($1, $2, $3, $4)`

	_, err := w.DB.Exec(query, req.City, req.Days)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (w *Weather) GetWeatherForecast(days int32, city string) ([]*pbt.GWFResponse, error) {
	forecasts := []*pbt.GWFResponse{}

	query := `select * from Weather where City = $1`

	rows, err := w.DB.Query(query, city, days)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		cur := pbt.GWFResponse{}
		err := rows.Scan(&cur.Time, &cur.Humidity, &cur.Temperature, &cur.WindMPH)
		if err != nil {
			return nil, err
		}
		forecasts = append(forecasts, &cur)
	}

	return forecasts, nil
}
