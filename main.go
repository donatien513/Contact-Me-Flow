package Handler

import "net/http"
import "github.com/julienschmidt/httprouter"
import "github.com/donatien513/Contact-Me-Flow/route"
import "github.com/donatien513/Contact-Me-Flow/utils"

  utils.InitRedisClient()
  Handler := httprouter.New()
  Handler.POST("/", route.HeartbeatHandler)
  Handler.POST("/ask-send", route.AuthentificationHandler)
