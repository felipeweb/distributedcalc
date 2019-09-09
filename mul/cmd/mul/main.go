package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/felipeweb/distributedcalc/mul"
	"github.com/felipeweb/distributedcalc/proto"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterMulServer(s, &mul.Server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
