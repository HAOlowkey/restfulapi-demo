package http

import (
	"github.com/HAOlowkey/restfulapi-demo/apps/host"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var Api = &handler{}

type handler struct {
	log  logger.Logger
	host host.ServiceServer
}

func (h *handler) Name() string {
	return host.AppName
}

func (h *handler) Config() error {
	h.log = zap.L().Named("HOST API")

	// if app.Host == nil {
	// 	panic("dependnece host service is nil")
	// }

	// h.host = app.Host
	h.host = app.GetGrpcApp(host.AppName).(host.ServiceServer)
	return nil
}

func (h *handler) Registry(r router.SubRouter) {
	r.Handle("POST", "/hosts", h.CreateHost)
	r.Handle("GET", "/hosts", h.QueryHost)
	r.Handle("GET", "/hosts/:id", h.DescribeHost)
	r.Handle("PUT", "/hosts/:id", h.PutUpdateHost)
	r.Handle("PATCH", "/hosts/:id", h.PatchUpdateHost)
	r.Handle("DELETE", "/hosts/:id", h.DeleteHost)

	// r.POST("/hosts", h.CreateHost)
	// // http.HandleFunc("/host", h.CreateHost)
	// r.GET("/hosts", h.QueryHost)
	// r.GET("/hosts/:id", h.DescribeHost)
	// r.PUT("/hosts/:id", h.PutUpdateHost)
	// r.PATCH("/hosts/:id", h.PatchUpdateHost)
	// r.DELETE("/hosts/:id", h.DeleteHost)
}

func init() {
	app.RegistryHttpApp(Api)
}
