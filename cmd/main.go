package main

import (
	"os"

	"github.mpi-internal.com/hu/homeoffice-calendar-api/cmd/handlers"
	"github.mpi-internal.com/hu/homeoffice-calendar-api/internal/config"
	"github.mpi-internal.com/hu/homeoffice-calendar-api/internal/logger"
	"github.mpi-internal.com/hu/homeoffice-calendar-api/internal/transport"
)

var (
	build   = "development"
	version = ""
)

func main() {
	cfg := config.New()

	l, err := logger.New(os.Stdout, "debug")
	if err != nil {
		panic(err)
	}

	l = l.With("build", build).With("version", version)

	handler := handlers.NewHandler(l)

	transport.Start(l, cfg.Transport, handler)
}
