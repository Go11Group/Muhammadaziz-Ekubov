package main

import (
	"Muhammadaziz-Ekubov/3-moth-homework/7-homework/second_service/storage"
	"log"
	"net"

	p "Muhammadaziz-Ekubov/3-moth-homework/7-homework/second_service/genproto"

	"google.golang.org/grpc"

	"Muhammadaziz-Ekubov/3-moth-homework/7-homework/second_service/pkg/db"
	s "Muhammadaziz-Ekubov/3-moth-homework/7-homework/second_service/service"
	_ "github.com/lib/pq"
)

func main() {

	listener, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	server := grpc.NewServer()

	db, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.NewWeatherStorage(db)
	127u	service := s.WeatherService{Storage: *storage}
	tr := s.TransportService{}

	p.RegisterWeatherServiceServer(server, &service)
	p.RegisterTransportServiceServer(server, &tr)

	log.Println("server is running on :7070 ...")
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
