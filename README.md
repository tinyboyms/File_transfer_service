# File Transfer Service

A high-performance file transfer service built with Go and gRPC, supporting efficient large file transfers through streaming.

## Features

- Bidirectional streaming for file transfers
- Chunked file upload and download
- File metadata retrieval
- Progress tracking
- Concurrent file operations

## Prerequisites

- Go 1.20 or later
- Protocol Buffers compiler
- gRPC tools

## Installation

1. Clone the repository

    git clone https://github.com/tinyboyms/File_transfer_service.git

2. Install dependencies
```bash
go mod download
 ```

## Usage

```bash
go mod download
 ```

## Usage
1. Start the server:
```bash
go run cmd/server/main.go
 ```

2. Use the client:
```bash
go run cmd/client/main.go
 ```

## API
The service provides three main operations:

- Upload files (client streaming)
- Download files (server streaming)
- Get file information (unary)
## Technologies
- Go
- gRPC
- Protocol Buffers
- Streaming Data Transfer    