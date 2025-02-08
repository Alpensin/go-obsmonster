package handlers

import (
	"encoding/json"
	"net/http"
)

type APIHandler func(w http.ResponseWriter, r *http.Request) error

type JSONResponse struct {
	Message string `json:"message,omitempty"`
}

func NewHandler(h APIHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			response, err := json.Marshal(JSONResponse{Message: err.Error()})
			if err != nil {
				w.Write([]byte("all broken"))
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		}
	}
}
