package main

import (
	"context"
	"log"
	"net"

	pb "github.com/yourusername/distributed-ml/pkg/communication"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMasterServiceServer
}

func (s *server) AssignTask(ctx context.Context, req *pb.TaskRequest) (*pb.TaskResponse, error) {
	// For simplicity, acknowledge the task assignment
	log.Printf("Assigned task: Algorithm=%s, Data=%v", req.Algorithm, req.Data)
	return &pb.TaskResponse{Message: "Task assigned"}, nil
}

func (s *server) ReceiveResult(ctx context.Context, req *pb.ResultRequest) (*pb.ResultResponse, error) {
	// Handle received result from worker
	log.Printf("Received result from %s: %v", req.WorkerId, req.Result)
	return &pb.ResultResponse{Message: "Result received"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMasterServiceServer(grpcServer, &server{})
	log.Println("Master server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
