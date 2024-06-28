package storage

import (
	"database/sql"

	genproto "Muhammadaziz-Ekubov/3-moth-homework/7-homework/second_service/genproto"
)

type WeatherStorage struct {
	db *sql.DB
}

func NewWeatherStorage(db *sql.DB) *WeatherStorage {
	return &WeatherStorage{
		db: db,
	}
}

func (u *WeatherStorage) GetCurrentWeather(req *genproto.Void) (*genproto.WeatherDaily, error) {

	//

	return &genproto.WeatherDaily{
		Date:     "2024-06-27",
		Location: "Tashkent",
		Weather: &genproto.WeatherCondition{
			Temperature: 32,
			Humidity:    "lightly",
			WindSpeed:   15,
			Condition:   "shiny",
		},
	}, nil
}

func (u *WeatherStorage) GetWeatherForecast(req *genproto.Date) (*genproto.WeatherData, error) {

	//

	return &genproto.WeatherData{
		Weather: []*genproto.WeatherDaily{
			{
				Date:     "2024-06-27",
				Location: "Tashkent",
				Weather: &genproto.WeatherCondition{
					Temperature: 32,
					Humidity:    "lightly",
					WindSpeed:   15,
					Condition:   "shiny",
				},
			},
			{
				Date:     "2024-06-27",
				Location: "Namangan",
				Weather: &genproto.WeatherCondition{
					Temperature: 32,
					Humidity:    "lightly",
					WindSpeed:   15,
					Condition:   "shiny",
				},
			},
		},
	}, nil
}

func (u *WeatherStorage) ReportWeatherCondition(req *genproto.WeatherDaily) (*genproto.Response, error) {

	//

	return nil, nil
}
