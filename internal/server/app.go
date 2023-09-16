package server

import (
	"net/http"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/nakiner/eth-parser/internal/logger"
	"github.com/oklog/run"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Port                    int           `envconfig:"SERVICE_PORT_HTTP" default:"8081"`
	GracefulShutdownTimeout time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT" default:"15s"`
	GracefulShutdownDelay   time.Duration `envconfig:"GRACEFUL_SHUTDOWN_DELAY" default:"30s"`
	LogLevel                zapcore.Level `envconfig:"LOG_LEVEL" default:"info"`
}

type App struct {
	runGroup run.Group
	http     *httpServer
	closer   *closer
}

var appLogger = logger.Logger()

func fromEnv() Config {
	var cfg Config
	envconfig.MustProcess("", &cfg)
	return cfg
}

func New() *App {
	appCfg := fromEnv()
	logger.SetLevel(appCfg.LogLevel)
	app := new(App)

	app.closer = new(closer)
	app.closer.add(func() error {
		return logger.Logger().Sync()
	})

	app.http = newHTTPServer(appCfg.Port, appCfg.GracefulShutdownTimeout)
	app.AddActor(interruptActor(appCfg.GracefulShutdownDelay))
	app.AddActor(app.http.actor())
	app.AddActor(app.closer.actor())

	return app
}

func (a *App) SetHandler(handler http.Handler) {
	a.http.server.Handler = handler
}

func (a *App) Use(fn func(next http.Handler) http.Handler) {
	a.http.server.Handler = fn(a.http.server.Handler)
}

func (a *App) Run() error {
	appLogger.Warn("application started")
	defer appLogger.Warn("application stopped")
	return a.runGroup.Run()
}

func (a *App) AddActor(execute func() error, interrupt func(error)) {
	a.runGroup.Add(execute, interrupt)
}

func (a *App) AddCloser(closer func() error) {
	a.closer.add(closer)
}
