package main

import (
	"github.com/Agile-Ultimate-English-team/foreign-at-home/internal/auth"
	pb "github.com/Agile-Ultimate-English-team/foreign-at-home/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func newGrpcServer(log *zap.SugaredLogger) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterAuthenticatorServer(s, auth.NewGrpcServer(log))
	return s
}
