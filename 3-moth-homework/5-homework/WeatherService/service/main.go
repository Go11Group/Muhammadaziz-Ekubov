package service

import (
	"context"
	"time"

	pb "Muhammadaziz-Ekubov/3-moth-homework/5-homework/WeatherService/genproto/weather_service"
	"Muhammadaziz-Ekubov/3-moth-homework/5-homework/WeatherService/storage/postgres"
)

type Weather struct {
	pb.UnimplementedWeatherServiceServer
	WeatherRepo *postgres.WeatherRepo
}

func NewWeather(weatherRepo *postgres.WeatherRepo) *Weather {
	return &Weather{WeatherRepo: weatherRepo}
}

func (w *Weather) GetCurrentWeather(ctx context.Context, req *pb.RequestCurrentWeather) (*pb.Weather, error) {
	currentWeather, err := w.WeatherRepo.GetCurrentWeather(req.City, time.Now())
	if err != nil {
		return &pb.Weather{}, err
	}
	return currentWeather, nil
}

func (w *Weather) GetWeatherForecast(ctx context.Context, req *pb.RequestWeatherForecast) (*pb.ResponceWeatherForecast, error) {
	forecast, err := w.WeatherRepo.GetWeatherForecast(req.City, time.Now(), int(req.Days))
	if err != nil {
		return &pb.ResponceWeatherForecast{}, err
	}
	return &pb.ResponceWeatherForecast{Forecastdays: *forecast}, err
}

func (w *Weather) ReportWeatherCondition(ctx context.Context, req *pb.Weather) (*pb.ResponceWeatherCondition, error) {
	err := w.WeatherRepo.CreateWeather(req)
	if err != nil {
		return &pb.ResponceWeatherCondition{Success: false}, err
	}
	return &pb.ResponceWeatherCondition{Success: true}, nil
}
