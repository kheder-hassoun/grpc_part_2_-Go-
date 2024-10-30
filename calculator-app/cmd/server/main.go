// cmd/server/main.go
package main

import (
    "context"
    "log"
    "net"
    "sync"

    "google.golang.org/grpc"
    "calculator-app/calculator"
)

// calculatorServer struct to hold server logic and history
type calculatorServer struct {
    calculator.UnimplementedCalculatorServiceServer
    history []calculator.HistoryEntry
    mu      sync.Mutex // To ensure thread-safe access to the history slice
}

// Add method implementation for addition service
func (s *calculatorServer) Add(ctx context.Context, req *calculator.TwoNumbers) (*calculator.CalculationResult, error) {
    result := req.Num1 + req.Num2
    res := &calculator.CalculationResult{Result: result}

    // Record the operation in history
    s.mu.Lock()
    s.history = append(s.history, calculator.HistoryEntry{
        Operation: "Addition",
        Num1:      req.Num1,
        Num2:      req.Num2,
        Result:    result,
    })
    s.mu.Unlock()

    return res, nil
}

// Multiply method implementation for multiplication service
func (s *calculatorServer) Multiply(ctx context.Context, req *calculator.TwoNumbers) (*calculator.CalculationResult, error) {
    result := req.Num1 * req.Num2
    res := &calculator.CalculationResult{Result: result}

    // Record the operation in history
    s.mu.Lock()
    s.history = append(s.history, calculator.HistoryEntry{
        Operation: "Multiplication",
        Num1:      req.Num1,
        Num2:      req.Num2,
        Result:    result,
    })
    s.mu.Unlock()

    return res, nil
}

// GetHistory method implementation for streaming history service
func (s *calculatorServer) GetHistory(req *calculator.HistoryRequest, stream calculator.CalculatorService_GetHistoryServer) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Stream each entry in the history to the client
    for _, entry := range s.history {
        if err := stream.Send(&entry); err != nil {
            return err
        }
    }
    return nil
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
