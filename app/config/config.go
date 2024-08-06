package config

import (
	"github.com/andresbott/go-carbon/libs/config"
	"strconv"
)

type AppCfg struct {
	Server serverCfg
	Obs    serverCfg `config:"Observability"`
	Auth   authConfig
	Log    log
	Msgs   []Msg
}

type log struct {
	Level string
}

type serverCfg struct {
	Disabled bool
	BindIp   string
	Port     int
}

func (c serverCfg) Addr() string {
	if c.BindIp == "" {
		return ":" + strconv.Itoa(c.Port)
	}
	return c.BindIp + ":" + strconv.Itoa(c.Port)
}

type authConfig struct {
	Enabled     bool // set to true to enable auth, default false
	SessionPath string
	HashKey     string
	BlockKey    string
	UserStore   userStore
}

type userStore struct {
	StoreType string `config:"Type"` // can be static | file
	FilePath  string `config:"Path"`
	Users     []User
}
type User struct {
	Name string
	Pw   string
}

// Default represents the basic set of sensible defaults
var defaultCfg = AppCfg{

	Server: serverCfg{
		BindIp: "",
		Port:   8085,
	},
	Obs: serverCfg{
		Disabled: true,
		BindIp:   "",
		Port:     9090,
	},
	Auth: authConfig{
		Enabled:     false, // set to true to enable auth, default false
		SessionPath: "",    // location where the sessions are stored
		HashKey:     "",    // cookie store encryption key
		BlockKey:    "",    // cookie value encryption
		UserStore: userStore{
			StoreType: "static",
			Users:     []User{},
		},
	},
	Log: log{
		Level: "info",
	},
}

type Msg struct {
	Level string
	Msg   string
}

const ConfigEnvVar = "FE26"

func Get(file string) (AppCfg, error) {
	configMsg := []Msg{}
	cfg := AppCfg{}
	var err error
	_, err = config.Load(
		config.Defaults{Item: defaultCfg},
		//config.CfgFile{Path: file}, // todo make file optional
		config.EnvVar{Prefix: ConfigEnvVar},
		config.Unmarshal{Item: &cfg},
		config.Writer{Fn: func(level, msg string) {
			if level == config.InfoLevel {
				configMsg = append(configMsg, Msg{Level: "info", Msg: msg})
			}
			if level == config.DebugLevel {
				configMsg = append(configMsg, Msg{Level: "debug", Msg: msg})
			}
		}},
	)
	cfg.Msgs = configMsg
	return cfg, err
}
