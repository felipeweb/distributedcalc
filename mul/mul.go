package mul

import (
	"context"
	"log"

	"github.com/felipeweb/distributedcalc/proto"
)

type Server struct {
}

func (s *Server) Mul(ctx context.Context, op *proto.OpRequest) (*proto.ResultResponse, error) {
	log.Println("GRPC Mul service")
	return &proto.ResultResponse{
		Result: op.GetLeft() * op.GetRight(),
	}, nil

}
