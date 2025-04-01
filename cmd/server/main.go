package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/yourusername/file_transfer/pkg/proto"
	"github.com/yourusername/file_transfer/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 50051, "The server port")
	uploadDir := flag.String("dir", "./uploads", "Directory to store uploaded files")
	flag.Parse()

	// Create server instance
	server, err := service.NewServer(*uploadDir)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	// Create listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterFileTransferServiceServer(grpcServer, server)

	// Handle graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		log.Println("Shutting down server...")
		grpcServer.GracefulStop()
	}()

	// Start server
	log.Printf("Server listening on port %d", *port)
	log.Printf("Upload directory: %s", *uploadDir)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}