package conf

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/labstack/gommon/log"
)

var (
	Conf              config // holds the global app config.
	defaultConfigFile = "conf/conf.toml"
)

type config struct {
	ReleaseMode bool   `toml:"release_mode"`
	LogLevel    string `toml:"log_level"`

	SessionStore string `toml:"session_store"`
	CacheStore   string `toml:"cache_store"`

	// 应用配置
	App app

	// 模板
	Tmpl tmpl

	Server server

	// MySQL
	DB database `toml:"database"`

	// 静态资源
	Static static

	// Redis
	Redis redis

	// Memcached
	Memcached memcached
}

type app struct {
	Name string `toml:"name"`
}

type server struct {
	Graceful bool   `toml:"graceful"`
	Addr     string `toml:"addr"`

	DomainApi    string `toml:"domain_api"`
	DomainWeb    string `toml:"domain_web"`
	DomainSocket string `toml:"domain_socket"`
}

type static struct {
	Type string `toml:"type"`
}

type tmpl struct {
	Type   string `toml:"type"`   // PONGO2,TEMPLATE(TEMPLATE Default)
	Data   string `toml:"data"`   // BINDATA,FILE(FILE Default)
	Dir    string `toml:"dir"`    // PONGO2(template/pongo2),TEMPLATE(template)
	Suffix string `toml:"suffix"` // .html,.tpl
}

type database struct {
	Name     string `toml:"name"`
	UserName string `toml:"user_name"`
	Pwd      string `toml:"pwd"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
}

type redis struct {
	Server string `toml:"server"`
	Pwd    string `toml:"pwd"`
}

type memcached struct {
	Server string `toml:"server"`
}

func init() {
}

// initConfig initializes the app configuration by first setting defaults,
// then overriding settings from the app config file, then overriding
// It returns an error if any.
func InitConfig(configFile string) error {
	if configFile == "" {
		configFile = defaultConfigFile
	}

	// Set defaults.
	Conf = config{
		ReleaseMode: false,
		LogLevel:    "DEBUG",
	}

	if _, err := os.Stat(configFile); err != nil {
		return errors.New("config file err:" + err.Error())
	} else {
		log.Infof("load config from file:" + configFile)
		configBytes, err := ioutil.ReadFile(configFile)
		if err != nil {
			return errors.New("config load err:" + err.Error())
		}
		_, err = toml.Decode(string(configBytes), &Conf)
		if err != nil {
			return errors.New("config decode err:" + err.Error())
		}
	}

	// @TODO 配置检查
	log.Infof("config data:%v", Conf)

	return nil
}

func GetLogLvl() log.Lvl {
	//DEBUG INFO WARN ERROR OFF
	switch Conf.LogLevel {
	case "DEBUG":
		return log.DEBUG
	case "INFO":
		return log.INFO
	case "WARN":
		return log.WARN
	case "ERROR":
		return log.ERROR
	case "OF":
		return log.OFF
	}

	return log.DEBUG
}

const (
	// Template Type
	PONGO2   = "PONGO2"
	TEMPLATE = "TEMPLATE"

	// Bindata
	BINDATA = "BINDATA"

	// File
	FILE = "FILE"

	// Redis
	REDIS = "REDIS"

	// Memcached
	MEMCACHED = "MEMCACHED"

	// Cookie
	COOKIE = "COOKIE"

	// In Memory
	IN_MEMORY = "IN_MEMARY"
)