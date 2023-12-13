package grpcserver

import (
	"context"
	"go_specs_greet/domain/interactions"
)

type Server struct {
	UnimplementedGreeterServer
}

func (g Server) Greet(ctx context.Context, request *GreetRequest) (*GreetReply, error) {
	return &GreetReply{Message: interactions.Greet(request.Name)}, nil
}

func (g Server) Curse(ctx context.Context, request *CurseRequest) (*CurseReplay, error) {
	return &CurseReplay{Message: interactions.Curse(request.Name)}, nil
}
