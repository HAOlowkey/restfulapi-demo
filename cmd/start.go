package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/HAOlowkey/restfulapi-demo/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/spf13/cobra"

	"github.com/HAOlowkey/restfulapi-demo/app"
	"github.com/HAOlowkey/restfulapi-demo/app/host/impl"
	"github.com/HAOlowkey/restfulapi-demo/protocol"
)

var (
	confFile, confType string
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long:  `start ...`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := loadGlobalConfig(confType); err != nil {
			return err
		}
		if err := loadGlobalLogger(); err != nil {
			return err
		}

		if err := impl.Service.Init(); err != nil {
			return err
		}

		app.Host = impl.Service

		svc := newService(conf.C())
		// 启动服务后，需要处理的事件
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)

		// 等待程序退出
		go svc.waitSign(ch)

		return svc.start()
	},
}

// http service
func newService(conf *conf.Config) *service {
	return &service{
		conf: conf,
		http: protocol.NewHttpService(),
		log:  zap.L().Named("Service"),
	}
}

type service struct {
	conf *conf.Config
	http *protocol.HttpService
	log  logger.Logger
}

func (s *service) start() error {
	return s.http.Start()
}

func (s *service) waitSign(sign chan os.Signal) {
	for sg := range sign {
		switch v := sg.(type) {
		// 留的扩展, 啥时候 需要单独处理时，再补充
		// 没做信号取反
		// reload
		// term
		// quick
		default:
			// 资源清理
			s.log.Infof("receive signal '%v', start graceful shutdown", v.String())
			if err := s.http.Stop(); err != nil {
				s.log.Errorf("graceful shutdown err: %s, force exit", err)
			}
			s.log.Infof("service stop complete")
			return
		}
	}
}

func loadGlobalConfig(configType string) error {
	switch configType {
	case "file":
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			return err
		}
	case "env":
		err := conf.LoadConfigFromEnv()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("not support config type")
	}
	return nil
}

func loadGlobalLogger() error {
	var (
		logInitMsg string
		level      zap.Level
	)

	// 获取出日志配置对象
	lc := conf.C().Log

	// debug, info, xxxx
	// 解析配置的日志级别是否正确
	lv, err := zap.NewLevel(lc.Level)
	if err != nil {
		// 解析失败, 默认使用Info级别
		logInitMsg = fmt.Sprintf("%s, use default level INFO", err)
		level = zap.InfoLevel
	} else {
		// 解析成功, 直接使用用户配置的日志级别
		level = lv
		logInitMsg = fmt.Sprintf("log level: %s", lv)
	}

	// 初始化了日志的默认配置
	zapConfig := zap.DefaultConfig()
	zapConfig.Level = level
	zapConfig.Files.RotateOnStartup = false
	switch lc.To {
	case conf.ToStdout:
		zapConfig.ToStderr = true
		zapConfig.ToFiles = false
	case conf.ToFile:
		zapConfig.Files.Name = "restful-api.log"
		zapConfig.Files.Path = lc.PathDir
	}
	switch lc.Format {
	case conf.JsonFormat:
		zapConfig.JSON = true
	}

	// 初始化全局Logger的配置
	if err := zap.Configure(zapConfig); err != nil {
		return err
	}

	// 全局Logger初始化后 就可以正常使用
	zap.L().Named("INIT").Info(logInitMsg)
	return nil
}

func init() {
	startCmd.PersistentFlags().StringVarP(&confFile, "conf_file", "f", "etc/restfulapi.toml", "the restful-api config file path")
	startCmd.PersistentFlags().StringVarP(&confType, "conf_type", "t", "file", "the restful-api config type")
	rootCmd.AddCommand(startCmd)
}
