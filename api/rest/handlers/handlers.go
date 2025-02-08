package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Alpensin/go-obsmonster/pkg/logging/console"
)

type APIHandler func(w http.ResponseWriter, r *http.Request) error

type JSONResponse struct {
	Message string `json:"message,omitempty"`
}

// NewHandler - simple handler wrapper for beginning
func NewHandler(logger console.Logger, h APIHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		err := h(w, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			response, err := json.Marshal(JSONResponse{Message: err.Error()})
			if err != nil {
				w.Write([]byte("all broken"))
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		}
		elapsedTime := time.Since(startTime)
		logger.Info("request",
			console.NewArg("method", r.Method),
			console.NewArg("url", r.URL.Path),
			console.NewArg("elapsedTime", elapsedTime),
		)
	}
}
