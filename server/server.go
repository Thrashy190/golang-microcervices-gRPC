package server

import (
	"context"
	"github.com/Thrashy190/go/grpc/models"
	"github.com/Thrashy190/go/grpc/repository"
	"github.com/Thrashy190/go/grpc/studentpb"
)

type Server struct {
	repo repository.Repository
	studentpb.UnimplementedStudentServiceServer
}

func NewStudentServer(repo repository.Repository) *Server {
	return &Server{repo: repo}
}

func (s *Server) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	student, err := s.repo.GetStudent(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &studentpb.Student{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func (s *Server) SetStudent(ctx context.Context, req *studentpb.Student) (*studentpb.SetStudentResponse, error) {
	student := &models.Student{
		Id:   req.Id,
		Name: req.Name,
		Age:  req.Age,
	}
	if err := s.repo.SetStudent(ctx, student); err != nil {
		return nil, err
	}
	return &studentpb.SetStudentResponse{
		Id: student.Id,
	}, nil
}
