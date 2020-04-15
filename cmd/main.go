package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/harmony-one/harmony-tf/config"
	"github.com/harmony-one/harmony-tf/testcases"
)

func main() {
	// Force usage of Go's own DNS implementation
	os.Setenv("GODEBUG", "netdns=go")

	if config.Args.PprofPort > 0 {
		go func() {
			fmt.Println(http.ListenAndServe(fmt.Sprintf("localhost:%d", config.Args.PprofPort), nil))
		}()
	}

	if err := execute(); err != nil {
		log.Fatalln(err)
	}

	// When we run using pprof we don't want to immediately exit,
	// We want to manually exit so that we have time to capture heap dumps etc. after the test suite has executed
	if config.Args.PprofPort > 0 {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		os.Exit(1)
	}
}

func execute() error {
	config.Execute()

	basePath, err := filepath.Abs(config.Args.Path)
	if err != nil {
		return err
	}

	if err := config.Configure(basePath); err != nil {
		return err
	}

	if err := testcases.Execute(); err != nil {
		return err
	}

	return nil
}
