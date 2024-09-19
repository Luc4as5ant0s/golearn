package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"

	algopb "github.com/Luc4as5ant0s/golearn/pkg/algorithms"
	pb "github.com/Luc4as5ant0s/golearn/pkg/communication"
)

func main() {
	// Command-line flag for worker ID
	workerID := flag.String("worker-id", "worker-1", "Worker ID")
	flag.Parse()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMasterServiceClient(conn)

	// Simulate requesting a task from the master
	taskRequest := &pb.TaskRequest{
		Algorithm: "mean",
		Data:      []float64{1, 2, 3, 4, 5},
		Params:    map[string]string{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	taskResponse, err := client.AssignTask(ctx, taskRequest)
	if err != nil {
		log.Fatalf("Failed to assign task: %v", err)
	}
	log.Printf("Master response: %s", taskResponse.Message)

	// Select the algorithm based on the task
	var algorithm algopb.Algorithm

	switch taskRequest.Algorithm {
	case "mean":
		algorithm = &algopb.MeanAlgorithm{}
	default:
		log.Fatalf("Unknown algorithm: %s", taskRequest.Algorithm)
	}

	// Initialize and compute
	err = algorithm.Initialize(nil)
	if err != nil {
		log.Fatalf("Algorithm initialization failed: %v", err)
	}

	result, err := algorithm.Compute(taskRequest.Data)
	if err != nil {
		log.Fatalf("Algorithm computation failed: %v", err)
	}
	log.Printf("Computed result: %v", result)

	// Send result back to master
	resultRequest := &pb.ResultRequest{
		WorkerId: *workerID,
		Result:   result,
	}

	resultResponse, err := client.ReceiveResult(ctx, resultRequest)
	if err != nil {
		log.Fatalf("Failed to send result: %v", err)
	}
	log.Printf("Master response: %s", resultResponse.Message)
}
