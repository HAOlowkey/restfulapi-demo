package impl

import (
	"database/sql"

	"github.com/HAOlowkey/restfulapi-demo/apps/host"
	"github.com/HAOlowkey/restfulapi-demo/conf"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
)

var Service = &impl{}

type impl struct {
	log logger.Logger
	db  *sql.DB
	host.UnimplementedServiceServer
}

func (i *impl) Config() error {
	var err error
	i.db, err = conf.C().Mysql.GetDB()
	if err != nil {
		return err
	}
	i.log = zap.L().Named("Host")
	return nil
}

func (i *impl) Name() string {
	return host.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	host.RegisterServiceServer(server, Service)
}

func init() {
	app.RegistryGrpcApp(Service)
}
