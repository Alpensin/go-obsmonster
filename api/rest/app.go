package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Alpensin/go-obsmonster/api/rest/handlers"
)

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello/{name}", handlers.NewHandler(handlers.Hello))

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		err := http.ListenAndServe(":8080", mux)
		if err != nil {
			fmt.Println(err)
			cancel()
		}
		cancel()
	}()
	<-ctx.Done()
}
