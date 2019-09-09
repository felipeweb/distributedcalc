package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/felipeweb/distributedcalc/proto"
	"github.com/felipeweb/distributedcalc/quo"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterQuoServer(s, &quo.Server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
