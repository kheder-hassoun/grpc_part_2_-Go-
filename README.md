# CalculatorService - Go gRPC Service

This project provides a simple gRPC-based CalculatorService written in Go. The service offers two operations: addition and multiplication, taking two numbers as input and returning the result.

## Project Structure

```
CalculatorService/
├── calculator/              # Contains the Go service implementation
│   ├── calculator.proto     # gRPC service definition
│   ├── calculator.pb.go     # Generated Protobuf code
│   └── calculator_grpc.pb.go # Generated gRPC code
├── cmd/
│   └── server/              # Main entry point for running the server
│       └── main.go
├── go.mod                   # Go module file
└── README.md                # Project documentation
```

## Prerequisites

- **Go 1.17** or higher
- **Protocol Buffers Compiler (protoc)**
- **protoc-gen-go** and **protoc-gen-go-grpc** plugins

### Install Protocol Buffers Compiler (protoc)

1. Download the latest version of `protoc` from the [official site](https://github.com/protocolbuffers/protobuf/releases).
2. Ensure `protoc` is in your system PATH.

### Install Go gRPC and Protocol Buffers Plugins

Install the `protoc` plugins for Go:
```bash
# Install Go plugins for protoc
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Make sure these plugins are available in your `PATH` so `protoc` can locate them.

## Generating Go Code from `.proto` File

From the project root, generate the Go code for gRPC using the following command:

```bash
# Generate code from calculator.proto
$ protoc --go_out=. --go_opt=paths=source_relative \
         --go-grpc_out=. --go-grpc_opt=paths=source_relative \
         calculator/calculator.proto
```

This generates two files:
- `calculator.pb.go`: Contains Protobuf message definitions.
- `calculator_grpc.pb.go`: Contains the gRPC service definitions.

## Running the Calculator Service

1. **Build the Server**:
   - From the project root, build the Go server:
     ```bash
     $ go build -o bin/server cmd/server/main.go
     ```

2. **Run the Server**:
   - Start the server by running the following command:
     ```bash
     $ ./bin/server
     ```
   - By default, the server listens on port `50051`. You can modify the port in the `main.go` file if needed.

3. **Client Interaction**:
   - The server exposes two RPC methods:
     - **Add**: Adds two numbers and returns the sum.
     - **Multiply**: Multiplies two numbers and returns the product.

You can create a client using the generated `calculator_grpc.pb.go` code to test the server or use any gRPC client tool to send requests.

## Project Dependencies

The dependencies for this project are managed in `go.mod`. Key dependencies include:

- `google.golang.org/grpc`: Go package for gRPC
- `google.golang.org/protobuf`: Protobuf support for Go

Run `go mod tidy` if you add or remove dependencies.

## Additional Notes

- The service is designed to handle concurrent requests, so feel free to extend it as needed.
- You can add more RPC methods to `calculator.proto` and regenerate the code using `protoc`.
- Customize logging and error handling in `main.go` or the service implementation for better visibility in production.
