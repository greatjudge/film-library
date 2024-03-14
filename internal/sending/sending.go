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

func JSONMarshalAndSend(w http.ResponseWriter, obj any, status int) {
	serialized, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(serialized)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
