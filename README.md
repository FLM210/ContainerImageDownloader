# Container Image Downloader

A web service for downloading container images from any registry.

## Features

- Simple web interface for downloading container images
- Supports Docker and OCI archive formats
- Configurable listening port
- Automatic file naming based on image name and tag
- Clean and modern UI with real-time status updates
- No zombie processes (proper child process management)

## Installation

### Prerequisites
- Go 1.16 or higher
- Skopeo tool installed
- Docker (optional, for containerized deployment)

### From Source

1. Clone the repository:
```bash
git clone https://github.com/FLM210/containerimagedownload.git
cd containerimagedownload
```

2. Build the application:
```bash
go build -o containerimagedownload main.go
```

3. Run the application:
```bash
./containerimagedownload
```

### Using Docker

1. Build the Docker image:
```bash
docker build -t containerimagedownload .
```

2. Run the container:
```bash
docker run -p 8080:8080 -v /path/to/downloads:/root/containerimagedownload containerimagedownload
```

## Configuration

### Environment Variables

- `LISTENPORT`: Controls the port on which the service listens (default: 8080)

### Example with Custom Port

```bash
LISTENPORT=3000 ./containerimagedownload
```

## Usage

1. Access the web interface at `http://localhost:8080` (or your custom port)
2. Enter the container image name (e.g., `nginx:latest` or `registry.example.com/nginx:latest`)
3. Select the output format (Docker Archive or OCI Archive)
4. Click the "Download Image" button
5. Wait for the download to complete, the file will be automatically downloaded

## Dependencies

- Go standard library
- Skopeo (for image copying functionality)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.