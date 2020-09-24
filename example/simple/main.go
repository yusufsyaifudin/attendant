package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yusufsyaifudin/attendant"
)

func main() {
	srvConf := attendant.Config{
		EnableProfiling: false,
		ListenAddress:   8000,
		WriteTimeout:    3000 * time.Second,
		ReadTimeout:     3000 * time.Second,
		ZapLogger:       nil,
		OpenTracing:     nil,
	}

	srv := attendant.NewServer(srvConf)

	var errChan = make(chan error, 1)
	go func() {
		errChan <- srv.Start()
	}()

	var signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalChan:
		_, _ = fmt.Fprintf(os.Stdout, "exiting...\n")
		srv.Shutdown()

	case err := <-errChan:
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error API: %s\n", err.Error())
		}
	}
}
