package internal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/justcompile/midgard/web/internal/websocket"

	log "github.com/sirupsen/logrus"
)

type HTTPServer struct {
	address string
	server  *http.Server
	hub     *websocket.Hub
}

func (s *HTTPServer) Listen() error {
	log.Infof("Starting Web server listening on: %s", s.address)
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err.Error() != "http: Server closed" {
			log.Fatal(err)
		}
	}()

	go s.hub.Run()

	return nil
}

func (s *HTTPServer) Shutdown() error {
	log.Info("Web server shutting down")
	s.hub.Shutdown()

	if err := s.server.Close(); err != nil && err.Error() != "http: Server closed" {
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	return nil
}

func NewHTTPServer(address string) *HTTPServer {
	hub := websocket.NewHub()
	router := NewRouter()

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(hub, w, r)
	})

	server := &HTTPServer{
		address: address,
		hub:     hub,
		server: &http.Server{
			Handler:      router,
			Addr:         address,
			WriteTimeout: 666 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
	}

	return server
}
