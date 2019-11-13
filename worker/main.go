package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/justcompile/midgard/common/workercomms"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

// printFeature gets the feature for the given point.
func connect(client pb.MidgardWorkerClient, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	connection, err := client.Connect(ctx, &pb.Worker{Name: name})
	if err != nil {
		log.Fatalf("%v.Connect(_) = _, %v: ", client, err)
	}
	log.Println(connection)
}

// printFeatures lists all the features within the given bounding Rectangle.
func disconnect(client pb.MidgardWorkerClient, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.Disconnect(ctx, &pb.Worker{Name: name})
	if err != nil {
		log.Fatalf("%v.Disconnect(_) = _, %v", client, err)
	}

	log.Println("Disconnected")
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	signals := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	client := pb.NewMidgardWorkerClient(conn)

	name, _ := os.Hostname()

	connect(client, name)

	go func() {
		<-signals
		fmt.Println()

		log.Println("Closing")
		disconnect(client, name)
		conn.Close()

		done <- true
	}()

	defer conn.Close()

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(30 * time.Second)
	ticker.Stop()
	done <- true
}
