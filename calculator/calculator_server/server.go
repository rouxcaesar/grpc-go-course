package main

import (
  "fmt"
  "context"
  "log"
  "net"
  "google.golang.org/grpc"
  "grpc-go-course/calculator/calculatorpb"
)

type server struct {}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
  fmt.Printf("Received Sum RPC: %v\n", req)
  firstNumber := req.FirstNumber
  secondNumber := req.SecondNumber
  sum := firstNumber + secondNumber
  res := &calculatorpb.SumResponse {
    SumResult: sum,
  }
  return res, nil
}

func main() {
  fmt.Println("Calculator Server")

  listener, err := net.Listen("tcp", "0.0.0.0:50051")
  if err != nil {
    log.Fatalf("Failed to listen: %v", err)
  }

  s := grpc.NewServer()
  calculatorpb.RegisterCalculatorServiceServer(s, &server{}) 

  if err := s.Serve(listener); err != nil {
    log.Fatalf("Failed to server: %v", err)
  }
}
