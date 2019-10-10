package main

import "net/http"
import "github.com/gorilla/mux"
import "github.com/donatien513/Contact-Me-Flow/route"
import "github.com/donatien513/Contact-Me-Flow/utils"

func main() {
  utils.InitRedisClient()
  router := mux.NewRouter()
  router.HandleFunc("/", route.HeartbeatHandler).Methods("GET")
  router.HandleFunc("/ask-send", route.AuthentificationHandler).Methods("POST")
  //router.HandleFunc("/confirm", route.confirmHandler).Methods("GET")
  http.Handle("/", router)
  http.ListenAndServe(":3000", nil)
}
