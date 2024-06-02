package main

import (
	"add/storage/CRUD"
	"add/storage/postgres"
	"fmt"
	"log"
)

func main() {
	db, err := postgres.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create user and product
	err = CRUD.CreateUserAndProduct(db, "grom", "jgssgg@example.com", "1455432", "Product1", "Description for Product1", 100.00, 10)
	if err != nil {
		log.Fatalf("CreateUserAndProduct failed: %v", err)
	}
	fmt.Println("User and Product created successfully")

	// Update user and product
	err = CRUD.UpdateUserAndProduct(db, 1, "john_doe_updated", "john_updated@example.com", "newsecurepassword", 1, "Product1_updated", "Updated description for Product1", 150.00, 20)
	if err != nil {
		log.Fatalf("UpdateUserAndProduct failed: %v", err)
	}
	fmt.Println("User and Product updated successfully")

	// Delete user and product
	err = CRUD.DeleteUserAndProduct(db, 1, 1)
	if err != nil {
		log.Fatalf("DeleteUserAndProduct failed: %v", err)
	}
	fmt.Println("User and Product deleted successfully")
}
