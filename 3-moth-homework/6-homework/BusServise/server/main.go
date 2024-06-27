package main

import (
	"Muhammadaziz-Ekubov/3-moth-homework/5-homework/BusServise/service"
	"Muhammadaziz-Ekubov/3-moth-homework/5-homework/BusServise/storage/postgres"
	"fmt"
	"log"
	"net"

	pb "Muhammadaziz-Ekubov/3-moth-homework/5-homework/BusServise/genproto/transport_service"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is listening on port 50051...")

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	s := service.NewBus(postgres.NewBusRepo(db))
	grpc := grpc.NewServer()
	pb.RegisterTransportServiceServer(grpc, s)

	if err := grpc.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
