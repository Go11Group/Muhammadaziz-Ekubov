package main

import (
	"fmt"
	"log"

	"context"
	pb "library/genprotos/generator"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	c := pb.NewLibraryServiceClient(conn)

	// // AddBook
	// req := &pb.AddBookRequest{
	// 	Title:         "Saodat Asri",
	// 	Author:        "Ahmad Lutfiy",
	// 	YearPublished: 2024,
	// }
	// r, err := c.AddBook(context.Background(), req)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(r)

	// // SearchBook
	// query := `select * from book`
	// reqS := &pb.SearchBookRequest{Query: query}
	// books, err := c.SearchBook(context.Background(), reqS)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(books)

	// Borrow book
	reqB := &pb.BorrowBookRequest{
		BookId: "3a7b38ac-3e11-4a21-ad55-518aba99e37d",
		UserId: uuid.NewString(),
	}
	resB, err := c.BorrowBook(context.Background(), reqB)
	if err != nil {
		log.Fatal(resB)
	}
	fmt.Println(resB)
}
