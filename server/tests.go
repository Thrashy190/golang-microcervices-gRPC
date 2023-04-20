package server

import (
	"context"
	"github.com/Thrashy190/go/grpc/models"
	"github.com/Thrashy190/go/grpc/repository"
	"github.com/Thrashy190/go/grpc/testpb"
)

type TestServer struct {
	repo repository.Repository
	testpb.UnimplementedTestServiceServer
}

func NewTestServer(repo repository.Repository) *TestServer {
	return &TestServer{repo: repo}
}

func (s *TestServer) GetTest(ctx context.Context, req *testpb.GetTestRequest) (*testpb.Test, error) {
	test, err := s.repo.GetTest(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &testpb.Test{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) SetTest(ctx context.Context, req *testpb.Test) (*testpb.SetTestResponse, error) {
	test := &models.Test{
		Id:   req.Id,
		Name: req.Name,
	}
	if err := s.repo.SetTest(ctx, test); err != nil {
		return nil, err
	}
	return &testpb.SetTestResponse{
		Id: test.Id,
	}, nil
}
