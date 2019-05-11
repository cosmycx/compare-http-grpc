package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Starting Validate RPC server...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	RegisterValidatePBServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} // .err

} // .main

func (*server) Validate(stream ValidatePBService_ValidateServer) error {
	fmt.Println("Validate invoked with streaming request ...")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// finished reading the streaming client
			return stream.SendAndClose(&ValidationPBResponse{
				Result: result,
			})
		} // .if
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		particip := req.GetParticipID()
		result += "particip: " + particip + ", "
	}
} // .Validate
