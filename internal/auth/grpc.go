package auth

import (
	"context"

	pb "github.com/Agile-Ultimate-English-team/foreign-at-home/proto"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewGrpcServer(log *zap.SugaredLogger) pb.AuthenticatorServer {
	return &authServer{
		log: log,
	}
}

type authServer struct {
	pb.UnimplementedAuthenticatorServer
	log *zap.SugaredLogger
}

func (s *authServer) AuthenticateByOAuth2Code(ctx context.Context, req *pb.OAuth2CodeAuthenticationRequest) (*emptypb.Empty, error) {
	s.log.Debugf("Authenticate by OAath2Code \"%s\" for %s provider", req.Code, req.Provider)
	panic("TODO")
	//	return &emptypb.Empty{}, nil
}
