package main

import (
	"context"
	"database/sql"
	pb "library/genprotos/generator"
	"library/server/storage/postgres"
	"log"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type LibraryService struct {
	pb.UnimplementedLibraryServiceServer
	DB *sql.DB
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(listener)
	}

	db, err := postgres.CreateDB()
	if err != nil {
		log.Fatal(err)
	}

	s := LibraryService{DB: db}
	grpc := grpc.NewServer()
	pb.RegisterLibraryServiceServer(grpc, &s)

	if err := grpc.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func (l *LibraryService) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	tx, err := l.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	newId := uuid.NewString()
	query := `insert into book(book_id, title, author, year_published)
	values ($1, $2, $3, $4)`

	_, err = tx.Exec(query, newId, req.Title, req.Author, req.YearPublished)
	if err != nil {
		return nil, err
	}
	return &pb.AddBookResponse{BookId: newId}, nil
}

func (l *LibraryService) SearchBook(ctx context.Context, req *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	res := []*pb.Book{}

	rows, err := l.DB.Query(req.Query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		book := pb.Book{}
		err := rows.Scan(&book.BookId, &book.Title, &book.Author,
			&book.YearPublished)

		if err != nil {
			return nil, err
		}

		res = append(res, &book)
	}
	return &pb.SearchBookResponse{Books: res}, nil
}

func (l *LibraryService) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	tx, err := l.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	query := `insert into borrowing(book_id, user_id) values($1, $2)`
	_, err = tx.Exec(query, req.BookId, req.UserId)
	if err != nil {
		return &pb.BorrowBookResponse{Success: false}, err
	}
	return &pb.BorrowBookResponse{Success: true}, nil
}
