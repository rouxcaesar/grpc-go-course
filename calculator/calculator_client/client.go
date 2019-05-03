package main

import (
  "context"
  "fmt"
  "google.golang.org/grpc"
  "grpc-go-course/calculator/calculatorpb"
  "log"
)

func main() {
  fmt.Println("Calculator Client")
  cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("Could not connect: %v", err)
  }
  defer cc.Close()

  c := calculatorpb.NewCalculatorServiceClient(cc)

  doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
  fmt.Println("Start to do a Sum Unary RPC...")
  req := &calculatorpb.SumRequest {
      FirstNumber: 5,
      SecondNumber: 40,
  }
  res, err := c.Sum(context.Background(), req)
  if err != nil {
    log.Fatalf("Error while calling Sum RPC: %v", err)
  }
  log.Printf("Response from Sum: %v", res.SumResult)
}
