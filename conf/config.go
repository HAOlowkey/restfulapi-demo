package conf

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
)

var (
	db     *sql.DB
	global *Config
)

func C() *Config {
	if global == nil {
		panic("config required")
	}
	return global
}

func newDefaultConfig() *Config {
	return &Config{
		Mysql: newDefaultMysql(),
		App:   newDefaultApp(),
		Log:   newDefaultLog(),
	}
}

func newDefaultMysql() *mysql {
	return &mysql{
		UserName:    "root",
		Password:    "123456",
		Port:        "3306",
		Database:    "api_demo",
		MaxOpenConn: 100,
		MaxIdleConn: 20,
		MaxLifeTime: 60 * 10,
		MaxIdleTime: 60 * 5,
	}
}

func newDefaultApp() *app {
	return &app{
		Name: "resetful-api",
		Host: "127.0.0.1",
		Port: "8090",
		Key:  "default app key",
	}
}

func newDefaultLog() *log {
	return &log{
		Level:  zap.DebugLevel.String(),
		To:     ToStdout,
		Format: TextFormat,
	}
}

func setGlobalConfig(conf *Config) {
	global = conf
}

type Config struct {
	Mysql *mysql
	Log   *log
	App   *app
}

type mysql struct {
	UserName string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Host     string `toml:"host" env:"MYSQL_HOST"`
	Port     string `toml:"port" env:"MYSQL_PORT"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`

	lock        sync.Mutex
	MaxOpenConn int `toml:"max_open_conn"`
	MaxIdleConn int `toml:"max_idel_conn"`
	MaxLifeTime int `toml:"max_life_time"`
	MaxIdleTime int `toml:"max_idle_time"`
}

func (m *mysql) GetDB() (*sql.DB, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if db == nil {
		conn, err := m.getDBConn()
		if err != nil {
			return nil, err
		}
		db = conn
	}
	return db, nil
}

func (m *mysql) getDBConn() (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", m.UserName, m.Password, m.Host, m.Port, m.Database)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}

type app struct {
	Name string `toml:"name"`
	Host string `toml:"host"`
	Port string `toml:"port"`
	Key  string `toml:"key"`
}

type log struct {
	Level   string    `toml:"level"`
	PathDir string    `toml:"dir"`
	Format  LogFormat `toml:"format"`
	To      LogTo     `toml:"to"`
}
