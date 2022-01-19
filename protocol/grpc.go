package protocol

import (
	"net"

	app "github.com/HAOlowkey/restfulapi-demo/apps"
	"github.com/HAOlowkey/restfulapi-demo/apps/host"
	"github.com/HAOlowkey/restfulapi-demo/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
)

func NewGrpcService() *GrpcService {
	server := grpc.NewServer()
	return &GrpcService{
		l: zap.L().Named("Grpc Service"),
		s: server,
	}
}

type GrpcService struct {
	l logger.Logger
	s *grpc.Server
}

func (g *GrpcService) Start() {
	host.RegisterServiceServer(g.s, app.Host)
	addr := conf.C().App.GrpcAddr()
	lsr, err := net.Listen("tcp", addr)
	if err != nil {
		g.l.Errorf("listen grpc tcp conn error, %s", err)
		return
	}
	g.l.Infof("GRPC 服务监听地址: %s", conf.C().App.GrpcAddr())
	if err := g.s.Serve(lsr); err != nil {
		if err == grpc.ErrServerStopped {
			g.l.Info("service is stopped")
		}

		g.l.Error("start grpc service error, %s", err.Error())
		return
	}
}

func (g *GrpcService) Stop() {
	g.s.GracefulStop()
}
