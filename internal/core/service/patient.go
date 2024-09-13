package service

import (
	"context"

	pb "github.com/sugaml/patient/internal/patient" // Adjust to your generated protobuf package
)

type Server struct {
	pb.UnimplementedPatientServiceServer
}

// GetPatient implements the logic for the GetPatient RPC
func (s *Server) GetPatient(ctx context.Context, req *pb.PatientRequest) (*pb.PatientResponse, error) {
	// Mocked response for demonstration purposes
	response := &pb.PatientResponse{
		Name: "John Doe",
		Age:  30,
	}
	return response, nil
}
