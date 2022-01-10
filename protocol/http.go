package protocol

import (
	"context"
	"fmt"
	"net/http"
	"time"

	hostAPI "github.com/HAOlowkey/restfulapi-demo/app/host/http"
	"github.com/HAOlowkey/restfulapi-demo/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/julienschmidt/httprouter"
)

func NewHttpService() *HttpService {
	r := httprouter.New()
	return &HttpService{
		r: r,
		l: zap.L().Named("Http Service"),
		server: &http.Server{
			Addr:              conf.C().App.Addr(),
			Handler:           r,
			ReadHeaderTimeout: 60 * time.Second,
			// 连接, client -->  server
			// 请求 1G, 60s读取完, 就超时
			ReadTimeout: 60 * time.Second,
			// 服务端详情数据超时
			// resp 1G, 60s读取不完
			WriteTimeout: 60 * time.Second,
			// http tcp 复用
			IdleTimeout: 60 * time.Second,
			//  header大小控制
			MaxHeaderBytes: 1 << 20, // 1M
		},
	}
}

type HttpService struct {
	server *http.Server
	r      *httprouter.Router
	l      logger.Logger
}

func (s *HttpService) Start() error {
	hostAPI.Api.Init()
	hostAPI.Api.Register(s.r)
	s.l.Infof("HTTP服务启动成功, 监听地址: %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info("service is stopped")
			return nil
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}
	return nil
}

func (s *HttpService) Stop() error {
	s.l.Info("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// 优雅关闭HTTP服务
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Errorf("graceful shutdown timeout, force exit")
	}
	return nil
}
