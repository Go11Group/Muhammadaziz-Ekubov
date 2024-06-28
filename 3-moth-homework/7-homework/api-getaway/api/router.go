package api

import (
	pb "Muhammadaziz-Ekubov/3-moth-homework/7-homework/api-getaway/genproto/transport_service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50013", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb.NewTransportServiceClient(conn)
	resp, err := c.TrackBusLocation(context.Background(), &pb.BusLocationRequest{BusNumber: "111"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
