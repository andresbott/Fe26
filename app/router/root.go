package router

import (
	"github.com/andresbott/Fe26/app/handlers/fileserver"
	"github.com/andresbott/Fe26/app/metainfo"
	"github.com/andresbott/Fe26/app/spa"
	"github.com/andresbott/go-carbon/app/handlrs"
	"github.com/andresbott/go-carbon/libs/auth"
	"github.com/andresbott/go-carbon/libs/http/handlers"
	"github.com/andresbott/go-carbon/libs/http/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/spf13/afero"
	"net/http"
	"path/filepath"
	"time"
)

type MyAppHandler struct {
	router *mux.Router
}

func (h *MyAppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

type AppCfg struct {
	Logger      *zerolog.Logger
	AuthEnabled bool
	AuthMngr    *auth.SessionMgr
	Users       auth.UserLogin
}

type authMeta struct {
	AuthEnabled bool `json:"enabled"`
}

// NewAppHandler generates the main url router handler to be used in the server
func NewAppHandler(cfg AppCfg) (*MyAppHandler, error) {

	r := mux.NewRouter()

	// add observability
	hist := middleware.NewHistogram("", nil, nil)
	r.Use(func(handler http.Handler) http.Handler {
		return middleware.PromLogMiddleware(handler, hist, cfg.Logger)
	})

	// TODO not to have in production
	throttle := middleware.ReqThrottle{
		MinDelay: 1500 * time.Millisecond,
		MaxDelay: 3000 * time.Millisecond,
		On:       false,
	}
	r.Use(throttle.Throttle)

	// todo this should reflect prod vs non-prod property
	genericErrorMessage := false

	// todo read the path from config
	absPath, err := filepath.Abs("./")
	if err != nil {
		return nil, err
	}
	fileServer := fileserver.FileServer(afero.NewBasePathFs(afero.NewOsFs(), absPath), "/api/v0/fs/")

	// this sub router does NOT enforce authentication
	apiRoute := r.PathPrefix("/api/v0").Subrouter()
	apiRoute.Use(func(handler http.Handler) http.Handler {
		return middleware.JsonErrMiddleware(handler, genericErrorMessage)
	})

	// add the info endpoint
	err = infoApi(apiRoute, cfg.AuthEnabled)
	if err != nil {
		return nil, err
	}

	if cfg.AuthEnabled {
		// add users handling to api
		userApi(apiRoute, cfg.AuthMngr, cfg.Users)

		// this sub router does enforce authentication
		protected := r.PathPrefix("/api/v0").Subrouter()
		protected.Use(func(handler http.Handler) http.Handler {
			return middleware.JsonErrMiddleware(handler, genericErrorMessage)
		}, cfg.AuthMngr.Middleware)
		apiRoute = protected
	}

	apiRoute.PathPrefix("/fs").Methods(http.MethodGet, http.MethodDelete, http.MethodPut, http.MethodPost).Handler(fileServer)
	apiRoute.PathPrefix("/fs").Handler(handlers.StatusErr(http.StatusMethodNotAllowed))

	// attach spa handler
	// if you want to serve the spa from the root, pass "/" to the spa handler and the path prefix
	// not that the SPA base and route needs to be adjusted accordingly
	spaHandler, err := spa.NewCarbonSpa("/")
	if err != nil {
		return nil, err
	}
	r.Methods(http.MethodGet).PathPrefix("/").Handler(spaHandler)

	return &MyAppHandler{
		router: r,
	}, nil
}

func infoApi(apiRoute *mux.Router, authEnabled bool) error {
	data := map[string]any{}
	data["auth"] = authMeta{AuthEnabled: authEnabled}

	type Info struct {
		Version   string `json:"version"`
		BuildTime string `json:"buildtime"`
		Commit    string `json:"commit"`
	}
	data["meta"] = Info{
		Version:   metainfo.Version,
		BuildTime: metainfo.BuildTime,
		Commit:    metainfo.ShaVer,
	}
	infoHndl, err := handlers.StaticInfo(data)
	if err != nil {
		return err
	}

	apiRoute.Path("/info").Methods(http.MethodGet).Handler(infoHndl)
	return nil
}

func userApi(apiRoute *mux.Router, session *auth.SessionMgr, users auth.UserLogin) {

	//  LOGIN
	apiRoute.Path("/user/login").Methods(http.MethodPost).Handler(handlrs.UserLoginHandler(session, users))

	apiRoute.Path("/user/login").Methods(http.MethodOptions).Handler(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

	}))
	apiRoute.Path("/user/login").Handler(handlers.StatusErr(http.StatusMethodNotAllowed))

	// LOGOUT
	apiRoute.Path("/user/logout").Handler(handlrs.UserLogoutHandler(session))

	// STATUS
	apiRoute.Path("/user/status").Methods(http.MethodGet).Handler(handlrs.UserStatusHandler(session))
	apiRoute.Path("/user/status").Handler(handlers.StatusErr(http.StatusMethodNotAllowed))

	// OPTIONS
	apiRoute.Path("/user/options").Methods(http.MethodGet).Handler(handlers.StatusErr(http.StatusNotImplemented))
	apiRoute.Path("/user/options").Handler(handlers.StatusErr(http.StatusMethodNotAllowed))
}
