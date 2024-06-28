package api

import (
	"Muhammadaziz-Ekubov/3-moth-homework/7-homework/api-getaway/api/handler"

	pbt "Muhammadaziz-Ekubov/3-moth-homework/7-homework/api-getaway/genproto/weather_service"

	pbw "Muhammadaziz-Ekubov/3-moth-homework/7-homework/api-getaway/genproto/transport_service"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func NewRouter(conn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()

	weather := pbt.NewWeatherServiceClient(conn)
	transport := pbw.NewTransportServiceClient(conn)

	handler := handler.NewHandler(weather, transport)

	router.GET("/get", handler.Get)

	return router
}
