package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Hello - returns hello json message. But if name is Error it return error
func Hello(w http.ResponseWriter, r *http.Request) error {
	name := r.PathValue("name")
	if name == "Error" {
		return errors.New("You asked to return error")
	}

	response, err := json.Marshal(JSONResponse{Message: fmt.Sprintf("Hello, %s", name)})
	if err != nil {
		return err
	}

	_, err = w.Write(response)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	return nil
}
