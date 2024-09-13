package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/sugaml/patient/internal/patient" // Adjust to your generated protobuf package

	"google.golang.org/grpc"
)

// server is used to implement patient.PatientServiceServer.
type server struct {
	pb.UnimplementedPatientServiceServer
}

// GetPatient implements the logic for the GetPatient RPC
func (s *server) GetPatient(ctx context.Context, req *pb.PatientRequest) (*pb.PatientResponse, error) {
	// Mocked response for demonstration purposes
	response := &pb.PatientResponse{
		Name: "John Doe",
		Age:  30,
	}
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPatientServiceServer(s, &server{})

	fmt.Println("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
