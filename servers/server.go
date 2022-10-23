package servers

import (
	"context"
	pb "github.com/bo-er/monthly/proto/company"
)

type Server struct {
	pb.UnimplementedDepartmentServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) CreateDepartment(ctx context.Context, in *pb.CreateDepartmentRequest) (*pb.Department, error) {
	return &pb.Department{Name: in.Department.Name, DepartmentId: in.Department.DepartmentId}, nil
}

func (s *Server) GetDepartment(ctx context.Context, in *pb.GetDepartmentRequest) (*pb.Department, error) {
	return &pb.Department{Name: in.Name, DepartmentId: in.DepartmentId}, nil
}
