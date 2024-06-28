package service

import (
	"context"
	"fmt"
	p "github.com/Go11Group/at_lesson/lesson47/second_service/genproto"
	"github.com/Go11Group/at_lesson/lesson47/second_service/storage"

	// grpcclinet "github.com/Go11Group/at_lesson/lesson47service/grpc-clinet"

	"log"
)

type WeatherService struct {
	p.UnimplementedWeatherServiceServer
	Storage storage.WeatherStorage
	// Clinets grpcclinet.ServiceMeneger
}

func (u *WeatherService) GetCurrentWeather(ctx context.Context, req *p.Void) (*p.WeatherDaily, error) {
	res, err := u.Storage.GetCurrentWeather(req)
	if err != nil {
		log.Println("error while creating user : ", err)
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}

func (u *WeatherService) GetWeatherForecast(ctx context.Context, req *p.Date) (*p.WeatherData, error) {
	res, err := u.Storage.GetWeatherForecast(req)
	if err != nil {
		log.Println("error while creating user : ", err)
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}

func (u *WeatherService) ReportWeatherCondition(ctx context.Context, req *p.WeatherDaily) (*p.Response, error) {
	res, err := u.Storage.ReportWeatherCondition(req)
	if err != nil {
		log.Println("error while creating user : ", err)
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}
