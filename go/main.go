package main

import (
	"context"
	pb "github.com/loveavoider/testGRPC/tree/main/go/proto"
)


type routServer struct {
}

func (s *routServer) Greetings(ctx context.Context, person *pb.Person) {

}