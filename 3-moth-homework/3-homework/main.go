package main

import (
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/all")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	client := http.Client{}

	req, err := http.NewRequest("POST", "url", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client.Do(req)
}
