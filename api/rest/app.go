package rest

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Alpensin/go-obsmonster/api/rest/handlers"
	"github.com/Alpensin/go-obsmonster/pkg/logging/console"
)

const (
	ServerAddres             = ":8080"
	GracefulShutdownDuration = 10 * time.Second
)

func Run(logger console.Logger) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello/{name}", handlers.NewHandler(logger, handlers.Hello))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logger.Critical("server serving", console.NewArg("error", err))
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), GracefulShutdownDuration)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Critical("server shutdown", console.NewArg("error", err))
	}
}
