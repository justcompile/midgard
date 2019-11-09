package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/justcompile/midgard/web/internal"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	// for _, path := range plugins.Registry {
	// 	path.Greet()
	// }
	signals := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	app := internal.NewApp("0.0.0.0:8000")

	go func() {
		<-signals
		fmt.Println()
		if err := app.Shutdown(); err != nil {
			os.Exit(1)
		}
		done <- true
	}()

	app.Start()

	<-done
	log.Info("Farewell")
}
