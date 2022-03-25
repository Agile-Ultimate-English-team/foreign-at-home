package main

import (
	"context"
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

	// App context
	appCtx, stopApp := context.WithCancel(context.Background())
	_ = appCtx
	log.Info("Successfully started the application, though it doesn't do anything for now. Press Ctrl+C to exit")

	// Graceful Shutdown
	wg := &sync.WaitGroup{}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	wg.Add(1)
	go func() {
		select {
		case <-sigCh:
			stopApp()
			log.Info("Shutting down...")
		}
		wg.Done()
	}()
	wg.Wait()
	log.Info("Goodbye!")
}
