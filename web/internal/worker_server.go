package internal

import (
	"net"

	"google.golang.org/grpc"

	"github.com/justcompile/midgard/common"
	pb "github.com/justcompile/midgard/common/workercomms"
	log "github.com/sirupsen/logrus"
)

type WorkerServer struct {
	address string
	server  *grpc.Server
}

func (w *WorkerServer) Listen() error {
	log.Infof("Starting Worker server listening on: %s", w.address)
	listener, err := net.Listen("tcp", w.address)
	if err != nil {
		return err
	}
	var opts []grpc.ServerOption
	w.server = grpc.NewServer(opts...)
	pb.RegisterMidgardWorkerServer(w.server, common.WorkerRegistry)

	go func() {
		if err := w.server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

func (w *WorkerServer) Shutdown() error {
	log.Info("Worker server shutting down")
	w.server.GracefulStop()

	return nil
}

func NewWorkerServer(address string) *WorkerServer {
	return &WorkerServer{
		address: address,
	}
}
