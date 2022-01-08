package impl

import (
	"database/sql"

	"github.com/HAOlowkey/restfulapi-demo/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var Service = &impl{}

type impl struct {
	log logger.Logger
	db  *sql.DB
}

func (i *impl) Init() error {
	var err error
	i.db, err = conf.C().Mysql.GetDB()
	if err != nil {
		return err
	}
	i.log = zap.L().Named("Host")
	return nil
}
