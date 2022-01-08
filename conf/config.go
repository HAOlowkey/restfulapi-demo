package conf

var global *Config

func C() *Config {
	if global == nil {
		panic("config required")
	}
	return global
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
