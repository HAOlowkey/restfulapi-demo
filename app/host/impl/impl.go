package impl

import (
	"database/sql"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var Service = &impl{}

type impl struct {
	log logger.Logger
	db  sql.DB
}

func (i *impl) Init() error {
	i.log = zap.L().Named("Host")
	return nil
}
