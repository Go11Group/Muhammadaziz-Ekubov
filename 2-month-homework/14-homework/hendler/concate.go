package handler

import (
	"fmt"
	"net/http"
)

func concatination(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Query()["name"][0] + r.URL.Query()["lastname"][0]))
	fmt.Println(r.URL.Query()["age"][0] + r.URL.Query()["age"][1])
}
