package main

import (
	"fmt"
	"log"
	"net"

	pb "weather_and_bus/weather/genproto/weather"
	"weather_and_bus/weather/service"
	"weather_and_bus/weather/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is listening on port 50050")

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	c := service.NewWeather(postgres.NewWeatherRepo(db))
	grpc := grpc.NewServer()
	pb.RegisterWeatherServiceServer(grpc, c)

	if err := grpc.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
