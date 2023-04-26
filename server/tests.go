package server

import (
	"context"
	"github.com/Thrashy190/go/grpc/models"
	"github.com/Thrashy190/go/grpc/repository"
	"github.com/Thrashy190/go/grpc/studentpb"
	"github.com/Thrashy190/go/grpc/testpb"
	"io"
	"log"
	"time"
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

func (s *TestServer) SetQuestions(stream testpb.TestService_SetQuestionsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}
		if err != nil {
			log.Fatalf("Error reading stream: %v", err)
			return err
		}
		question := &models.Question{
			Id:       msg.GetId(),
			Answer:   msg.GetAnswer(),
			Question: msg.GetQuestion(),
			TestId:   msg.GetTestId(),
		}
		err = s.repo.SetQuestion(context.Background(), question)
		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}

	}
}

func (s *TestServer) EnrollStudents(stream testpb.TestService_EnrollStudentsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetEnrollmentResponse{
				Ok: true,
			})
		}
		if err != nil {
			log.Fatalf("Error reading stream: %v", err)
			return err
		}
		enrollment := &models.Enrollment{
			TestId:    msg.GetTestId(),
			StudentId: msg.GetStudentId(),
		}
		err = s.repo.SetEnrollment(context.Background(), enrollment)
		if err != nil {
			return stream.SendAndClose(&testpb.SetEnrollmentResponse{
				Ok: false,
			})
		}

	}
}

func (s *TestServer) GetStudentsPerTest(req *testpb.GetStudentsPerTestRequest, stream testpb.TestService_GetStudentsPerTestServer) error {
	students, err := s.repo.GetStudentsPerTest(context.Background(), req.GetTestId())
	if err != nil {
		return err
	}
	for _, student := range students {
		student := &studentpb.Student{
			Id:   student.Id,
			Name: student.Name,
			Age:  student.Age,
		}

		if err := stream.Send(student); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
