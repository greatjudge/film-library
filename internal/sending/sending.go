package sending

import (
	"encoding/json"
	"net/http"
)

func JSONError(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	err = json.NewEncoder(w).Encode(err)
	if err != nil {
		// TODO add log?
		w.WriteHeader(http.StatusBadRequest)
	}
}
