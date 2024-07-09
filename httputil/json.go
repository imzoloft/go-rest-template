package httputil

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, p any) error {
	if r.Body == nil || r.ContentLength == 0 {
		fmt.Print("no body")
		return errors.New("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(p)
}

func WriteJSON(w http.ResponseWriter, status int, p any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(p)
}

func WriteError(w http.ResponseWriter, status int, errors map[string]any) error {
	return WriteJSON(w, status, errors)
}
