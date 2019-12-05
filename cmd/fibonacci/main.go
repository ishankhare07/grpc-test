package main

import (
	"context"
	"grpc-test/pkg/services/fibonacci"
	"log"
	"net"

	"google.golang.org/grpc"
)

var port = ":8888"

func fibN(n int32) int32 {
	if n <= 1 {
		return n
	}

	return fibN(n-1) + fibN(n-2)
}

type FibServer struct{}

func (f *FibServer) GetNthFib(ctx context.Context, nth *fibonacci.FibNthRequest) (*fibonacci.FibNthResponse, error) {
	nthNum := fibN(nth.N)
	return &fibonacci.FibNthResponse{
		Value: nthNum,
	}, nil
}

func (f *FibServer) GetFibSeq(ctx context.Context, limit *fibonacci.FibSeqUpperLimit) (*fibonacci.FibSeqResponse, error) {
	seq := []int32{}
	var t1, t2 int32 = 0, 1

	for ; limit.UpperBound > 0; limit.UpperBound-- {
		seq = append(seq, t1)
		next := t1 + t2
		t1 = t2
		t2 = next
	}

	return &fibonacci.FibSeqResponse{
		Seq: seq,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf(err.Error())
	}

	s := grpc.NewServer()

	fibonacci.RegisterFibonacciServer(s, &FibServer{})
	log.Printf("Server listening on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf(err.Error())
	}
}
