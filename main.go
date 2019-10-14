package main

import "net/http"
import "github.com/gorilla/mux"
import "github.com/donatien513/Contact-Me-Flow/route"
import "github.com/donatien513/Contact-Me-Flow/utils"

func Handler() http.HandlerFunc {
  utils.InitRedisClient()
  router := mux.NewRouter()
  router.HandleFunc("/", route.HeartbeatHandler).Methods("POST")
  router.HandleFunc("/ask-send", route.AuthentificationHandler).Methods("POST")
  return http.Handle("/", router)
}
