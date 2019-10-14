package Handler

import "net/http"
import "github.com/julienschmidt/httprouter"
import "github.com/donatien513/Contact-Me-Flow/route"
import "github.com/donatien513/Contact-Me-Flow/utils"

func Handler() http.HandlerFunc {
  utils.InitRedisClient()
  router := httprouter.New()
  router.POST("/", route.HeartbeatHandler)
  router.POST("/ask-send", route.AuthentificationHandler)
  return router
}
