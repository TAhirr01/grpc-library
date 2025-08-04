package main

import (
	"context"
	"fmt"
	"github.com/TAhirr01/grpc-library/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	// 1.Connect to gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewLibraryServiceClient(conn)
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 2.Register user
	//userRegisterRes, err := client.RegisterUser(ctx, &pb.RegisterUserRequest{
	//	Name:  "Test",
	//	Email: "test1@example.com",
	//})
	//if err != nil {
	//	log.Fatalf("Error registering user: %v", err)
	//}
	//log.Printf("Registered User: ID=%d ,Name=%s", userRegisterRes.Id, userRegisterRes.Name)
	//
	//// 3.Add book
	//addBookRes, err := client.AddBook(ctx, &pb.AddBookRequest{Title: "testTitle", Author: "testAuthor"})
	//if err != nil {
	//	log.Fatalf("Error creating a user: %v", err)
	//}
	//log.Printf("Book created Book: ID=%d ,Title=%s", addBookRes.Id, addBookRes.Title)
	//
	//// 4.Add book to user
	//if _, err := client.BorrowBook(ctx, &pb.BorrowBookRequest{UserId: userRegisterRes.Id, BookId: addBookRes.Id}); err != nil {
	//	log.Fatalf("Error borrowing a book: %v", err)
	//}
	//log.Printf("Book borrowed Book: ID=%d ,User: ID=%d", addBookRes.Id, userRegisterRes.Id)
	// 4.List user's books
	//bookList, err := client.ListUserBooks(ctx, &pb.UserID{UserId: userRegisterRes.Id})
	//if err != nil {
	//	log.Fatalf("Error listing books: %v", err)
	//}
	//log.Println("User's Books")
	//for _, book := range bookList.Books {
	//	log.Printf(" - %s bys %s", book.Title, book.Author)
	//}

	//5.Upload books with stream
	stream, err := client.UploadBooks(context.Background())
	if err != nil {
		log.Fatalf("Error uploading books: %v", err)
	}

	books := []pb.AddBookRequest{
		{Title: "Go in Action", Author: "Caleb Doxsey"},
		{Title: "Clean Code", Author: "Robert C. Martin"},
		{Title: "Sehv kitab"},
		{Title: "Domain-Driven Design", Author: "Eric Evans"},
	}
	// Send books one by one
	for _, book := range books {
		if err := stream.Send(&book); err != nil {
			log.Fatalf("Failed to send book: %v", err)
		}
		fmt.Printf("Sent book: %s\n", book.Title)
		time.Sleep(500 * time.Millisecond)
	}

	// Close stream and receive summary
	summary, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf(" Upload failed: %v", err)
	}

	fmt.Printf(" Uploaded %d books successfully!\n", summary.Count)
}
