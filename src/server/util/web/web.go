package web

import (
	"encoding/json"
	"net/http"
)

// WriteResp : Better than simply fmt.Fprintf
func WriteResp(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	w.Write([]byte(msg))
}

// NewMessage : get a json message
func NewMessage(message string) []byte {
	var m struct {
		Message string `json:"message"`
	}
	m.Message = message

	json, _ := json.Marshal(m)
	return json
}

// SendJSON : send something under json
func SendJSON(w http.ResponseWriter, code int, content []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(content)
}
