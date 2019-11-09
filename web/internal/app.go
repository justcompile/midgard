package internal

import (
	log "github.com/sirupsen/logrus"
)

type App struct {
	webServer    Server
	workerServer Server
}

func (a *App) Start() {
	log.Info("Starting Midgard...")
	a.webServer.Listen()
	a.workerServer.Listen()
}

func (a *App) Shutdown() error {
	err := a.webServer.Shutdown()

	if err != nil {
		log.Error(err)
	}

	err = a.workerServer.Shutdown()
	if err != nil {
		log.Error(err)
	}

	return err
}

func NewApp(addr string) *App {
	return &App{
		webServer:    NewHTTPServer(addr),
		workerServer: NewWorkerServer("0.0.0.0:10000"),
	}
}
