package respone

import (
	"encoding/json"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, status int, object interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(object)
}
