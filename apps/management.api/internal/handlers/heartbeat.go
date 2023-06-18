package handlers

import "net/http"

func heartbeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
