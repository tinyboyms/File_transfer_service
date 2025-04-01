package service

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	pb "github.com/yourusername/file_transfer/pkg/proto"
	"google.golang.org/grpc"
)

// Client handles file transfer operations
type Client struct {
	client pb.FileTransferServiceClient
	conn   *grpc.ClientConn
}

// NewClient creates a new file transfer client
func NewClient(serverAddr string) (*Client, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %w", err)
	}

	return &Client{
		client: pb.NewFileTransferServiceClient(conn),
		conn:   conn,
	}, nil
}

// Close closes the client connection
func (c *Client) Close() error {
	return c.conn.Close()
}

// UploadFile uploads a file to the server
func (c *Client) UploadFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	fileSize := fileInfo.Size()
	filename := filepath.Base(filePath)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	stream, err := c.client.UploadFile(ctx)
	if err != nil {
		return fmt.Errorf("failed to create upload stream: %w", err)
	}

	buffer := make([]byte, 64*1024) // 64KB chunks
	var offset int64 = 0
	startTime := time.Now()

	for {
		bytesRead, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}

		chunk := &pb.FileChunk{
			Filename:  filename,
			Content:   buffer[:bytesRead],
			TotalSize: fileSize,
			ChunkSize: int64(bytesRead),
			Offset:    offset,
		}

		if err := stream.Send(chunk); err != nil {
			return fmt.Errorf("failed to send chunk: %w", err)
		}

		offset += int64(bytesRead)
		progress := float64(offset) / float64(fileSize) * 100
		elapsed := time.Since(startTime).Seconds()
		speed := float64(offset) / elapsed / 1024 / 1024 // MB/s

		fmt.Printf("\rUploading %s: %.2f%% (%d/%d bytes) - %.2f MB/s", 
			filename, progress, offset, fileSize, speed)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("failed to receive upload response: %w", err)
	}

	elapsed := time.Since(startTime).Seconds()
	speed := float64(fileSize) / elapsed / 1024 / 1024 // MB/s

	fmt.Printf("\nUpload complete: %s\n", resp.Message)
	fmt.Printf("Transferred %d bytes in %.2f seconds (%.2f MB/s)\n", 
		fileSize, elapsed, speed)

	return nil
}

// DownloadFile downloads a file from the server
func (c *Client) DownloadFile(filename, savePath string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	req := &pb.FileRequest{
		Filename: filename,
	}

	stream, err := c.client.DownloadFile(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to create download stream: %w", err)
	}

	// Create the directory if it doesn't exist
	saveDir := filepath.Dir(savePath)
	if err := os.MkdirAll(saveDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	file, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	var fileSize int64
	var bytesReceived int64
	startTime := time.Now()

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error receiving chunk: %w", err)
		}

		if fileSize == 0 {
			fileSize = chunk.TotalSize
		}

		_, err = file.WriteAt(chunk.Content, chunk.Offset)
		if err != nil {
			return fmt.Errorf("failed to write chunk: %w", err)
		}

		bytesReceived += chunk.ChunkSize
		progress := float64(bytesReceived) / float64(fileSize) * 100
		elapsed := time.Since(startTime).Seconds()
		speed := float64(bytesReceived) / elapsed / 1024 / 1024 // MB/s

		fmt.Printf("\rDownloading %s: %.2f%% (%d/%d bytes) - %.2f MB/s", 
			filename, progress, bytesReceived, fileSize, speed)
	}

	elapsed := time.Since(startTime).Seconds()
	speed := float64(fileSize) / elapsed / 1024 / 1024 // MB/s

	fmt.Printf("\nDownload complete: %s\n", filename)
	fmt.Printf("Transferred %d bytes in %.2f seconds (%.2f MB/s)\n", 
		fileSize, elapsed, speed)

	return nil
}

// GetFileInfo retrieves information about a file
func (c *Client) GetFileInfo(filename string) (*pb.FileInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := &pb.FileRequest{
		Filename: filename,
	}

	return c.client.GetFileInfo(ctx, req)
}