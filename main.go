package Handler

import "net/http"
import "github.com/donatien513/Contact-Me-Flow/routes"
import "github.com/donatien513/Contact-Me-Flow/utils"

func Handler(w http.ResponseWriter, r *http.Request) {
  utils.InitRedisClient()
  utils.InitRecaptcha()
  if r.URL.Path == "/ask-send" && r.Method == http.MethodPost {
    routes.AuthentificationHandler(w, r)
    return
  }
  if r.URL.Path == "/send" && r.Method == http.MethodPost {
    routes.DirectSendHandler(w, r)
    return
  }
  http.NotFound(w, r)
}
