// cmd/server/main.go
package main

import (
    "context"
    "log"
    "net"
    "google.golang.org/grpc"
    "calculator-app/calculator" // Import local calculator package
)

type calculatorServer struct {
    calculator.UnimplementedCalculatorServiceServer
}

func (s *calculatorServer) Add(ctx context.Context, req *calculator.TwoNumbers) (*calculator.CalculationResult, error) {
    result := req.Num1 + req.Num2
    return &calculator.CalculationResult{Result: result}, nil
}

func (s *calculatorServer) Multiply(ctx context.Context, req *calculator.TwoNumbers) (*calculator.CalculationResult, error) {
    result := req.Num1 * req.Num2
    return &calculator.CalculationResult{Result: result}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    calculator.RegisterCalculatorServiceServer(s, &calculatorServer{})

    log.Printf("Server is listening on port 50051...")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
