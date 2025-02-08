package main

import (
	"github.com/Alpensin/go-obsmonster/api/rest"
	"github.com/Alpensin/go-obsmonster/pkg/logging/console/mux"
)

func main() {
	logMux := mux.New()
	logMux.Info("start")
	rest.Run(logMux)
	logMux.Info("stop")
}
