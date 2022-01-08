package http

import (
	"github.com/HAOlowkey/restfulapi-demo/app"
	"github.com/HAOlowkey/restfulapi-demo/app/host"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/julienschmidt/httprouter"
)

var api = handler{}

type handler struct {
	log  logger.Logger
	host host.Service
}

func (h *handler) Init() {
	h.log = zap.L().Named("HOST API")

	if app.Host == nil {
		panic("dependnece host service is nil")
	}

	h.host = app.Host
}

func (h *handler) Register(r httprouter.Router) {
	r.POST("/host", h.CreateHost)
	// http.HandleFunc("/host", h.CreateHost)
	r.GET("/host", h.QueryHost)

}
