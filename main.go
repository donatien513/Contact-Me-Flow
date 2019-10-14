package Handler

import "net/http"
import "github.com/donatien513/Contact-Me-Flow/route"
import "github.com/donatien513/Contact-Me-Flow/utils"

func Handler(w http.ResponseWriter, r *http.Request) {
  utils.InitRedisClient()
  if r.URL.Path == "/" && r.Method == http.MethodGet {
    route.HeartbeatHandler(w, r)
    return
  }
  if r.URL.Path == "/ask-send" && r.Method == http.MethodPost {
    route.AuthentificationHandler(w, r)
    return
  }
  http.NotFound(w, r)
}
