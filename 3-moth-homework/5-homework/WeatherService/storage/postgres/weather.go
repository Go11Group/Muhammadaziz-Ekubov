package postgres

import (
	"database/sql"
	"time"

	pb "Muhammadaziz-Ekubov/3-moth-homework/5-homework/WeatherService/genproto/weather_service"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type WeatherRepo struct {
	DB *sql.DB
}

func NewWeatherRepo(db *sql.DB) *WeatherRepo {
	return &WeatherRepo{DB: db}
}

func (w *WeatherRepo) CreateWeather(newWeather *pb.Weather) error {
	tx, err := w.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `insert into weathers(city, date, temperature, humidity, wind) 
	values ($1, $2, $3, $4, $5)`

	_, err = tx.Exec(query, newWeather.City, newWeather.Date.AsTime(),
		newWeather.Temperature, newWeather.Humidity, newWeather.Wind)
	return err
}

func (w *WeatherRepo) GetCurrentWeather(city string, date time.Time) (*pb.Weather, error) {
	currentWeather := pb.Weather{Date: timestamppb.New(date)}

	query := `select city, temperature, humidity, wind from weathers where 
	city = $1 and date = $2`
	err := w.DB.QueryRow(query, city, date).Scan(&currentWeather.City,
		&currentWeather.Temperature, &currentWeather.Humidity,
		&currentWeather.Wind)

	if err != nil {
		return &pb.Weather{}, err
	}
	return &currentWeather, nil
}

func (w *WeatherRepo) GetWeatherForecast(city string, date time.Time, days int) (*[]*pb.Weather, error) {
	res := []*pb.Weather{}

	query := `select * from weathers where 
	city = $1 and date >= $2 and date <= $3`
	rows, err := w.DB.Query(query, city, date, date.Add(time.Hour*24*time.Duration(days)))
	if err != nil {
		return &[]*pb.Weather{}, err
	}

	for rows.Next() {
		w := &pb.Weather{Date: timestamppb.New(date)}
		var date time.Time
		err := rows.Scan(&w.City, &date, &w.Temperature, &w.Humidity, &w.Wind)
		if err != nil {
			return &[]*pb.Weather{}, err
		}
		res = append(res, w)
	}
	return &res, nil
}
