package main

import (
	"net"
	"os"
	"os/signal"
	"sync"

	"go.uber.org/zap"
)

func main() {
	// Logging
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	log := logger.Sugar()
	log.Info("Initialized sugared logger.")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Can't listen on port %d: %+v", 8080, err)
	}
	grpcSrv := newGrpcServer(log)

	// Graceful Shutdown
	wg := &sync.WaitGroup{}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	wg.Add(1)
	go func() {
		select {
		case <-sigCh:
			grpcSrv.GracefulStop()
			log.Info("Shutting down...")
		}
		wg.Done()
	}()

	//Grpc server
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := grpcSrv.Serve(lis); err != nil {
			log.Fatalf("Server has suddenly stopped: %+v", err)
		}
	}()
	log.Infof("Server is listening for GRPC connections on %s", lis.Addr())

	wg.Wait()
	log.Info("Goodbye!")
}
