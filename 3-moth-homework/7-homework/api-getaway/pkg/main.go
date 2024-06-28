package pkg

import (
	pbt "Muhammadaziz-Ekubov/3-moth-homework/7-homework/api-getaway/genproto/transport_service"
	pbw "Muhammadaziz-Ekubov/3-moth-homework/7-homework/api-getaway/genproto/weather_service"

	"google.golang.org/grpc"
)

func CreateTransportServicClient(conn *grpc.ClientConn) *pbt.TransportServiceClient {
	trc := pbt.NewTransportServiceClient(conn)
	return &trc
}

func CreateWheatherServicClient(conn *grpc.ClientConn) *pbw.WeatherServiceClient {
	whc := pbw.NewWeatherServiceClient(conn)
	return &whc
}
