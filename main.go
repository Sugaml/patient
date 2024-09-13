package main

import (
	"fmt"
	"log"
	"net"

	"github.com/sugaml/patient/internal/core/service"
	pb "github.com/sugaml/patient/internal/patient" // Adjust to your generated protobuf package

	"google.golang.org/grpc"
)

// server is used to implement patient.PatientServiceServer.

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPatientServiceServer(s, &service.Server{})

	fmt.Println("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
