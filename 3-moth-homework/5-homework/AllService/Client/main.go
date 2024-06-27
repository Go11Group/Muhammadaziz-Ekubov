package main

import (
	"log"

	pbTransport "Muhammadaziz-Ekubov/3-moth-homework/5-homework/AllService/genproto/transport_service"
	pbWeather "Muhammadaziz-Ekubov/3-moth-homework/5-homework/AllService/genproto/weather_service"

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

func CLIWeather(weatherClient *pbWeather.WeatherServiceClient) {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// resp, err := (*transportClient).GetBusSchedule(ctx, &pbTransport.RequestBusSchedule{BusNumber: 5})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp.Schedules)
}

func CLITransport(transportClient *pbTransport.TransportServiceClient) {
	// // get bus schedule
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// resp, err := (*transportClient).GetBusSchedule(ctx, &pbTransport.RequestBusSchedule{BusNumber: 5})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp.Schedules)

	// // get current location of bus
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// resp, err := (*transportClient).TrackBusLocation(ctx, &pbTransport.RequestBusLocation{BusNumber: 5})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp.StationName)

	// // give report for traffic Jam
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// resp, err := (*transportClient).ReportTrafficJam(ctx, &pbTransport.RequestTrafficJam{Report: "ok"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp.Success)
}
