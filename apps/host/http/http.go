package http

import (
	app "github.com/HAOlowkey/restfulapi-demo/apps"
	"github.com/HAOlowkey/restfulapi-demo/apps/host"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/julienschmidt/httprouter"
)

var Api = handler{}

type handler struct {
	log  logger.Logger
	host host.ServiceServer
}

func (h *handler) Init() {
	h.log = zap.L().Named("HOST API")

	if app.Host == nil {
		panic("dependnece host service is nil")
	}

	h.host = app.Host
}

func (h *handler) Register(r *httprouter.Router) {
	r.POST("/hosts", h.CreateHost)
	// http.HandleFunc("/host", h.CreateHost)
	r.GET("/hosts", h.QueryHost)
	r.GET("/hosts/:id", h.DescribeHost)
	r.PUT("/hosts/:id", h.PutUpdateHost)
	r.PATCH("/hosts/:id", h.PatchUpdateHost)
	r.DELETE("/hosts/:id", h.DeleteHost)
}
