package main

import (
	"github.com/Thrashy190/go/grpc/database"
	"github.com/Thrashy190/go/grpc/server"
	"github.com/Thrashy190/go/grpc/studentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {

	list, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatalf("Error listening: %s", err.Error())
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	newServer := server.NewStudentServer(repo)

	if err != nil {
		log.Fatalf("Error creating repository: %s", err.Error())
	}

	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, newServer)

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatalf("Error serving: %s", err.Error())
	}
}
