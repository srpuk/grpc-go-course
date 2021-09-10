package main

import (
	"context"
	"fmt"
	"grpc-g-course/greet/calculator/calculatorpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer//why i am not sure
}

func(*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error){
     
	fmt.Printf("Got the request rpc : %v",req)
	first:=req.FirstNum
	second:=req.SecondNum

	 sum:=first+second

	 res:=&calculatorpb.SumResponse{
	 	Result: sum,
	 }
	 return res,nil
}


func main() {

	fmt.Println("Calculator Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatal("Failed to listen", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s,&server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}