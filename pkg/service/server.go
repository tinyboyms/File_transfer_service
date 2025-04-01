package service

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	pb "github.com/yourusername/file_transfer/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server implements the FileTransferService
type Server struct {
	pb.UnimplementedFileTransferServiceServer
	uploadDir string
	mu        sync.Mutex
}

// NewServer creates a new file transfer server
func NewServer(uploadDir string) (*Server, error) {
	// Create upload directory if it doesn't exist
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %w", err)
	}

	return &Server{
		uploadDir: uploadDir,
	}, nil
}

// UploadFile handles file uploads from clients
func (s *Server) UploadFile(stream pb.FileTransferService_UploadFileServer) error {
	var filename string
	var fileSize int64
	var file *os.File
	var bytesReceived int64

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			// End of file transfer
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "failed to receive chunk: %v", err)
		}

		// For the first chunk, create the file
		if file == nil {
			filename = chunk.Filename
			fileSize = chunk.TotalSize
			
			// Ensure the filename is safe
			safePath := filepath.Join(s.uploadDir, filepath.Base(filename))
			
			var err error
			file, err = os.Create(safePath)
			if err != nil {
				return status.Errorf(codes.Internal, "failed to create file: %v", err)
			}
			defer file.Close()
			
			fmt.Printf("Starting upload of %s (%d bytes)\n", filename, fileSize)
		}

		// Write the chunk to the file at the correct offset
		_, err = file.WriteAt(chunk.Content, chunk.Offset)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to write chunk: %v", err)
		}

		bytesReceived += chunk.ChunkSize
		progress := float64(bytesReceived) / float64(fileSize) * 100
		fmt.Printf("Upload progress: %.2f%% (%d/%d bytes)\n", progress, bytesReceived, fileSize)
	}

	return stream.SendAndClose(&pb.FileResponse{
		Message: fmt.Sprintf("Successfully uploaded %s", filename),
		Success: true,
		Size:    bytesReceived,
	})
}

// DownloadFile handles file downloads to clients
func (s *Server) DownloadFile(req *pb.FileRequest, stream pb.FileTransferService_DownloadFileServer) error {
	filename := req.Filename
	
	// Ensure the filename is safe
	safePath := filepath.Join(s.uploadDir, filepath.Base(filename))
	
	file, err := os.Open(safePath)
	if err != nil {
		return status.Errorf(codes.NotFound, "file not found: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return status.Errorf(codes.Internal, "failed to get file info: %v", err)
	}

	fileSize := fileInfo.Size()
	fmt.Printf("Starting download of %s (%d bytes)\n", filename, fileSize)

	buffer := make([]byte, 64*1024) // 64KB chunks
	var offset int64 = 0

	for {
		bytesRead, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "error reading file: %v", err)
		}

		chunk := &pb.FileChunk{
			Filename:  filename,
			Content:   buffer[:bytesRead],
			TotalSize: fileSize,
			ChunkSize: int64(bytesRead),
			Offset:    offset,
		}

		if err := stream.Send(chunk); err != nil {
			return status.Errorf(codes.Internal, "failed to send chunk: %v", err)
		}

		offset += int64(bytesRead)
		progress := float64(offset) / float64(fileSize) * 100
		fmt.Printf("Download progress: %.2f%% (%d/%d bytes)\n", progress, offset, fileSize)
	}

	fmt.Printf("Download of %s completed\n", filename)
	return nil
}

// GetFileInfo returns information about a file
func (s *Server) GetFileInfo(ctx context.Context, req *pb.FileRequest) (*pb.FileInfo, error) {
	filename := req.Filename
	
	// Ensure the filename is safe
	safePath := filepath.Join(s.uploadDir, filepath.Base(filename))
	
	fileInfo, err := os.Stat(safePath)
	if err != nil {
		if os.IsNotExist(err) {
			return &pb.FileInfo{
				Filename: filename,
				Size:     0,
				Exists:   false,
			}, nil
		}
		return nil, status.Errorf(codes.Internal, "failed to get file info: %v", err)
	}

	return &pb.FileInfo{
		Filename: filename,
		Size:     fileInfo.Size(),
		Exists:   true,
	}, nil
}