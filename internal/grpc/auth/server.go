package auth

import (
	"context"

	ssov1 "github.com/kirakulakov/gp_grpc_protos/gen/go/sso"
	"google.golang.org/grpc"
)

type serverApi struct {
	ssov1.UnimplementedAuthServer // для того
	// чтобы можно было запустить приложение без реализации всех необходимых методов
	// заявленных в интерфейсе AuthServer
	// например на этапе разработки

	// если убрать, то для того чтобы можно было запустить приложение,
	// необходимо будет реализовать все методы
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverApi{})
}

func (s *serverApi) Login(
	ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	panic("not implemented")
}

func (s *serverApi) Register(
	ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	panic("not implemented")
}

func (s *serverApi) IsAdmin(
	ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	panic("not implemented")
}
