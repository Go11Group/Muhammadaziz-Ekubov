package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "user=postgres password=1111 dbname=test sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "SELECT id, name, age FROM student WHERE age >= $1"

	rows, err := db.Query(query, 19)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %s, Name: %s, Age: %d\n", id, name, age)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
