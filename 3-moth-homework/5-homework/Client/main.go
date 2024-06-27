package main

import (
	"log"

	pbTransport "Muhammadaziz-Ekubov/3-moth-homework/5-homework/BusServise/genproto/transport_service"
	pbWeather "Muhammadaziz-Ekubov/3-moth-homework/5-homework/WeatherService/genproto/weather_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	transportClient := CreateTransportServerClient()
	weatherClient := CreateWeatherServerClient()

	CLITransport(transportClient)
	CLIWeather(weatherClient)
}

func CreateTransportServerClient() *pbTransport.TransportServiceClient {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	c := pbTransport.NewTransportServiceClient(conn)
	return &c
}

func CreateWeatherServerClient() *pbWeather.WeatherServiceClient {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	c := pbWeather.NewWeatherServiceClient(conn)
	return &c
}
