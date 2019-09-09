package add

import (
	"context"
	"log"

	"github.com/felipeweb/distributedcalc/proto"
)

type Server struct {
}

func (s *Server) Add(ctx context.Context, op *proto.OpRequest) (*proto.ResultResponse, error) {
	log.Println("GRPC Add service")
	return &proto.ResultResponse{
		Result: op.GetLeft() + op.GetRight(),
	}, nil

}
