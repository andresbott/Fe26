package cmd

import (
	"fmt"
	"github.com/andresbott/Fe26/app/config"
	"github.com/andresbott/Fe26/app/metainfo"
	"github.com/andresbott/Fe26/app/router"
	"github.com/andresbott/go-carbon/libs/auth"
	"github.com/andresbott/go-carbon/libs/http/handlers"
	"github.com/andresbott/go-carbon/libs/http/server"
	"github.com/andresbott/go-carbon/libs/logzero"
	"github.com/andresbott/go-carbon/libs/user"
)

func runServer(configFile string) error {

	cfg, err := config.Get(configFile)
	if err != nil {
		return err
	}

	// setup the logger
	logOutput, err := logzero.ConsoleFileOutput("")
	if err != nil {
		return err
	}
	l := logzero.DefaultLogger(logzero.GetLogLevel(cfg.Log.Level), logOutput)

	l.Info().Str("version", metainfo.Version).Str("component", "startup").
		Msgf("running version %s, build date: %s, commint: %s ", metainfo.Version, metainfo.BuildTime, metainfo.ShaVer)

	// print config messages delayed
	for _, m := range cfg.Msgs {
		if m.Level == "info" {
			l.Info().Str("component", "config").Msg(m.Msg)
		} else {
			l.Debug().Str("component", "config").Msg(m.Msg)
		}
	}

	// Main APApplication handler
	appCfg := router.AppCfg{
		Logger:      l,
		AuthEnabled: false,
	}

	if cfg.Auth.Enabled {
		appCfg.AuthEnabled = true
		// session based auth
		//cookieStore, err := auth.CookieStore(hashKey, blockKey)
		cookieStore, err := auth.FsStore(cfg.Auth.SessionPath, []byte(cfg.Auth.HashKey), []byte(cfg.Auth.BlockKey))
		if err != nil {
			return err
		}
		sessionAuth, err := auth.NewSessionMgr(auth.SessionCfg{
			Store: cookieStore,
		})
		if err != nil {
			return err
		}
		appCfg.AuthMngr = sessionAuth

		var users auth.UserLogin
		// load the correct user manager
		switch cfg.Auth.UserStore.StoreType {
		case "static":
			staticUsers := user.StaticUsers{}
			for _, u := range cfg.Auth.UserStore.Users {
				staticUsers.Add(u.Name, u.Pw)
			}
			users = &staticUsers
			l.Debug().Str("component", "users").Msgf("loading %d static user(s)", len(staticUsers.Users))
		case "file":
			if cfg.Auth.UserStore.FilePath == "" {
				return fmt.Errorf("no path for users file is empty")
			}
			staticUsers, err := user.FromFile(cfg.Auth.UserStore.FilePath)
			if err != nil {
				return err
			}
			users = staticUsers
			l.Debug().Str("component", "users").Msgf("loading %d users from file", len(staticUsers.Users))
		default:
			return fmt.Errorf("wrong user store in configuration, %s is not supported", cfg.Auth.UserStore.StoreType)
		}
		appCfg.Users = users

	}

	rootHandler, err := router.NewAppHandler(appCfg)
	if err != nil {
		return err
	}

	s, err := server.New(server.Cfg{
		Addr:       cfg.Server.Addr(),
		Handler:    rootHandler,
		SkipObs:    cfg.Obs.Disabled,
		ObsAddr:    cfg.Obs.Addr(),
		ObsHandler: handlers.Observability(),
		Logger: func(msg string, isErr bool) {
			if isErr {
				l.Warn().Str("component", "server").Msg(msg)
			} else {
				l.Info().Str("component", "server").Msg(msg)
			}
		},
	})
	if err != nil {
		return err
	}

	return s.Start()
}
