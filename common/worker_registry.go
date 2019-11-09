package common

import (
	"context"
	"sync"

	"github.com/jinzhu/copier"
	"github.com/justcompile/midgard/common/events"
	pb "github.com/justcompile/midgard/common/workercomms"
)

type workerRegistry struct {
	workers map[string]*pb.Worker
	mu      sync.Mutex
}

func (s *workerRegistry) GetConnectedWorkers() []pb.Worker {
	s.mu.Lock()
	results := make([]pb.Worker, len(s.workers))
	i := 0

	for _, worker := range s.workers {
		var newWorker pb.Worker

		copier.Copy(&newWorker, worker)
		results[i] = newWorker

		i++
	}

	s.mu.Unlock()
	return results
}

// Connect is called when a worker connects to the pool
func (s *workerRegistry) Connect(ctx context.Context, worker *pb.Worker) (*pb.Connection, error) {
	s.mu.Lock()
	s.workers[worker.Name] = worker
	s.mu.Unlock()

	if err := events.WorkerConnected(map[string]interface{}{"name": worker.Name}); err != nil {
		return nil, err
	}

	return &pb.Connection{State: 200}, nil
}

// Disconnect is called when a worker disconnects from the pool
func (s *workerRegistry) Disconnect(ctx context.Context, worker *pb.Worker) (*pb.Connection, error) {
	s.mu.Lock()

	delete(s.workers, worker.Name)

	s.mu.Unlock()

	if err := events.WorkerDisconnected(map[string]interface{}{"name": worker.Name}); err != nil {
		return nil, err
	}

	return &pb.Connection{State: 200}, nil
}

// WorkerRegistry ...
var WorkerRegistry *workerRegistry

func init() {
	WorkerRegistry = &workerRegistry{
		workers: make(map[string]*pb.Worker),
	}
}
