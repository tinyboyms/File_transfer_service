package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/yourusername/file_transfer/pkg/service"
)

func main() {
	serverAddr := flag.String("server", "localhost:50051", "The server address in the format host:port")
	upload := flag.String("upload", "", "Path to file to upload")
	download := flag.String("download", "", "Name of file to download")
	output := flag.String("output", "", "Path to save downloaded file")
	info := flag.String("info", "", "Get information about a file")
	flag.Parse()

	// Create client
	client, err := service.NewClient(*serverAddr)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Handle commands
	if *upload != "" {
		// Upload file
		fmt.Printf("Uploading file: %s\n", *upload)
		if err := client.UploadFile(*upload); err != nil {
			log.Fatalf("Failed to upload file: %v", err)
		}
	} else if *download != "" {
		// Download file
		outputPath := *output
		if outputPath == "" {
			// If no output path is specified, save to current directory
			outputPath = filepath.Base(*download)
		}
		fmt.Printf("Downloading file: %s to %s\n", *download, outputPath)
		if err := client.DownloadFile(*download, outputPath); err != nil {
			log.Fatalf("Failed to download file: %v", err)
		}
	} else if *info != "" {
		// Get file info
		fileInfo, err := client.GetFileInfo(*info)
		if err != nil {
			log.Fatalf("Failed to get file info: %v", err)
		}
		
		if fileInfo.Exists {
			fmt.Printf("File: %s\n", fileInfo.Filename)
			fmt.Printf("Size: %d bytes\n", fileInfo.Size)
			fmt.Printf("Status: Exists on server\n")
		} else {
			fmt.Printf("File: %s\n", fileInfo.Filename)
			fmt.Printf("Status: Does not exist on server\n")
		}
	} else {
		// No command specified
		fmt.Println("Please specify a command:")
		fmt.Println("  -upload <file>: Upload a file to the server")
		fmt.Println("  -download <file> [-output <path>]: Download a file from the server")
		fmt.Println("  -info <file>: Get information about a file on the server")
		os.Exit(1)
	}
}