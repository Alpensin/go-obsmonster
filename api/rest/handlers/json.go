package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Hello - returns hello json message.
// If PathValue name equals Error it returns error
// If PathValue name equals Panic it panics
func Hello(w http.ResponseWriter, r *http.Request) error {
	name := r.PathValue("name")
	if name == "Error" {
		return errors.New("You asked to return error")
	}

	if name == "Panic" {
		panic("You asked to panic")
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
