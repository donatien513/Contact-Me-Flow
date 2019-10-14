package Handler

import "github.com/julienschmidt/httprouter"
import "github.com/donatien513/Contact-Me-Flow/route"
import "github.com/donatien513/Contact-Me-Flow/utils"

utils.InitRedisClient()
var Handler := httprouter.New()
Handler.POST("/", route.HeartbeatHandler)
Handler.POST("/ask-send", route.AuthentificationHandler)
