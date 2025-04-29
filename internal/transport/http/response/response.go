package response

import (
	"encoding/json"
	"net/http"
)

func NewResponse(s any, status int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(s)
	return
}
