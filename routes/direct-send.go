package routes

import "net/http"
import "encoding/json"
import "github.com/donatien513/Contact-Me-Flow/utils"
import "github.com/donatien513/Contact-Me-Flow/types"
import "github.com/donatien513/Contact-Me-Flow/config"

func DirectSendHandler(w http.ResponseWriter, r *http.Request) {
  if r.Body == nil {
    httpFailure(w, http.StatusBadRequest)
    return
  }
  decoder := json.NewDecoder(r.Body)
  var emailData types.EmailPendingRequest
  decodeErr := decoder.Decode(&emailData)
  if decodeErr != nil {
    httpFailure(w, http.StatusBadRequest)
    return
  }

  if emailData.Sender == "" || emailData.Message == "" {
    httpFailure(w, http.StatusBadRequest)
    return
  }

  emailDeliveryData := types.EmailDelivery{}
  emailDeliveryData.Recipients = []string{config.MyEmailAddress}
  emailDeliveryData.Body = emailData.Sender + ":<br /><br />" + emailData.Message
  utils.RequestEmailDelivery(emailDeliveryData)
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "text/plain")
}

// Handle Http response failure
func httpFailure(w http.ResponseWriter, httpStatusCode int) {
  http.Error(w, http.StatusText(httpStatusCode), httpStatusCode)
}
