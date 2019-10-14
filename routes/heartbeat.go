package routes

import "net/http"

func HeartbeatHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "text/plain")
  w.Write([]byte("I am working :)"))
}
