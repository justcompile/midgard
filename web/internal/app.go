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
	if err := a.webServer.Listen(); err != nil {
		log.Errorf("%v", err)
	}

	if err := a.workerServer.Listen(); err != nil {
		log.Errorf("%v", err)
	}
}

func (a *App) Shutdown() error {
	var err error

	if err = a.webServer.Shutdown(); err != nil {
		log.Error(err)
	}

	if err = a.workerServer.Shutdown(); err != nil {
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
