package quo

import (
	"context"
	"errors"
	"log"

	"github.com/felipeweb/distributedcalc/proto"
)

type Server struct {
}

func (s *Server) Quo(ctx context.Context, op *proto.OpRequest) (*proto.ResultResponse, error) {
	log.Println("GRPC Quo service")
	if op.GetRight() == 0 {
		return nil, errors.New("0 can't be on the right side in the division")
	}
	return &proto.ResultResponse{
		Result: op.GetLeft() / op.GetRight(),
	}, nil

}
