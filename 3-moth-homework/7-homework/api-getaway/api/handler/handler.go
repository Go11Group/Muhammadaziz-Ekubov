package handler

import (
	pbt "Muhammadaziz-Ekubov/3-moth-homework/7-homework/api-getaway/genproto/transport_service"
	pbw "Muhammadaziz-Ekubov/3-moth-homework/7-homework/api-getaway/genproto/weather_service"
)

type Handler struct {
	TrClient pbt.TransportServiceClient
	WhClient pbw.WeatherServiceClient
}

func NewHandler(Weather pbw.WeatherServiceClient, Transport pbt.TransportServiceClient) *Handler {
	return &Handler{WhClient: Weather, TrClient: Transport}
}
