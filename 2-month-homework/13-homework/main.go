package main

import (
	"log"
	controller "my_project/API"
	"net/http"
)

func main() {
	mux := http.ServeMux{}
	controller.StudentAPI(&mux)
	controller.UniversityAPI(&mux)
	controller.GroupAPI(&mux)
	controller.FacultyAPI(&mux)

	err := http.ListenAndServe(":8088", &mux)
	if err != nil {
		log.Println("Server listening in port :8080 ... ")
	}

}
