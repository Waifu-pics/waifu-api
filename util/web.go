package util

import (
	"net/http"
)

// WriteResp : Better than simply fmt.Fprintf
func WriteResp(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	w.Write([]byte(msg))
}
