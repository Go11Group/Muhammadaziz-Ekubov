package main

import (
	client "client/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main(){
	r:=gin.Default()

	r.POST("/clientcreatuser",client.Hanleclinetcreateuser)
	r.PUT("/clientupdateuser",client.Handleclientupdateuser)
	r.DELETE("/clientdeleteuser",client.Handleclientdeleteuser)
	r.GET("/clientreatuser",client.HandleclientReadalluser)

	if err:=r.Run(":8081");err!=nil{
		log.Fatal("8081 portni ishlatishda xatolik?",err)
	}
}