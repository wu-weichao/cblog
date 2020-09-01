package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

// app config
type App struct {
	JwtSecret string
	JwtTtl    int
	PageSize  int
}

var AppSetting = &App{}

// server config
type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

// database config
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
	Charset     string
}

var DatabaseSetting = &Database{}

// log config
type Log struct {
	Path  string
	Level int
}

var LogSetting = &Log{}

func init() {
	cfg, err := ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("setting.init parse 'config/app.ini': %v", err)
	}

	// parse setting
	MapTo(cfg, "app", AppSetting)
	MapTo(cfg, "server", ServerSetting)
	MapTo(cfg, "database", DatabaseSetting)
	MapTo(cfg, "log", LogSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

func MapTo(cfg *ini.File, name string, conf interface{}) {
	err := cfg.Section(name).MapTo(conf)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", name, err)
	}
}
