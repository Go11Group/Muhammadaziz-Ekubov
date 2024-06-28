package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

type Student struct {
	Id         int
	First_name string
	Last_name  string
	Email      string
	Gender     string
}

type Students struct {
	Students []Student
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /Students", GetStudents)

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	f, err := os.OpenFile("Students.json", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(data)
	if err != nil {
		panic(err)
	}
}

// decoder := json.NewDecoder(f)
// students := Students{}

// err = decoder.Decode(&students)
// if err != nil {
// 	panic(err)
// }

// fmt.Println(students)
