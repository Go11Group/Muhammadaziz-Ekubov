package client

import (
	"bytes"
	"client/model"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hanleclinetcreateuser(g *gin.Context) {
	var user model.Users

	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userdata, err := json.Marshal(user)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.Post("http://localhost:8080/createuser", "application/json", bytes.NewBuffer(userdata))
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	byt, err := io.ReadAll(resp.Body)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	g.Data(200, "application/json", byt)
}


func Handleclientupdateuser(g *gin.Context){
	var user model.Users

	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userupdate, err := json.Marshal(user)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}

	clients:=http.Client{}
	resp, err := http.NewRequest(http.MethodPut,"http://localhost:8080/updateuser",bytes.NewBuffer(userupdate))
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	resp.Header.Set("Content-Type", "application/json")
	clients.Do(resp)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		g.JSON(500, gin.H{"error": "Javobni o'qishda xatolik"})
		return
	}

	g.Data(http.StatusOK, "application/json", body)
}

func Handleclientdeleteuser(g *gin.Context){
	var user model.Users

	if err:=g.ShouldBindJSON(&user);err!=nil{
		g.JSON(400,gin.H{"error":err.Error()})
		return
	}

	userdelete, err := json.Marshal(user)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}

	clients:=http.Client{}
	resp, err := http.NewRequest(http.MethodDelete,"http://localhost:8080/deleteuser",bytes.NewBuffer(userdelete))
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	resp.Header.Set("Content-Type", "application/json")
	clients.Do(resp)

	byt, err := io.ReadAll(resp.Body)
	if err != nil {
		g.JSON(500, gin.H{"error": "Javobni o'qishda xatolik"})
		return
	}

	g.Data(200, "application/json", byt)
}

func HandleclientReadalluser(g *gin.Context){
	resp,err:=http.Get("http://localhost:8080/readuser")
	if err!=nil{
		g.JSON(500,gin.H{"error":err.Error()})
		return
	}
	defer resp.Body.Close()

	byt,err:=io.ReadAll(resp.Body)
	if err!=nil{
		g.JSON(500,gin.H{"error":err.Error()})
		return
	}

	g.Data(200,"application/json",byt)
}