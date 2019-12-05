package main

import (
	"context"
	"grpc-test/pkg/services/fibonacci"
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

func (f *FibServer) GetFibSeq(_ context.Context, _ *fibonacci.FibSeqUpperLimit) (*fibonacci.FibSeqResponse, error) {
	panic("not implemented") // TODO: Implement
}
