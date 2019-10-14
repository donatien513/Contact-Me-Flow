package routes

import "time"
import "sync"
import "net/http"
import "encoding/json"
import "github.com/donatien513/Contact-Me-Flow/utils"
import "github.com/donatien513/Contact-Me-Flow/types"

func AuthentificationHandler(w http.ResponseWriter, r *http.Request) {
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

  var emailPendingKey string = utils.GenerateEmailPendingKey()
  var waitGroup sync.WaitGroup
  waitGroup.Add(2)
  go func () {
    utils.StorePendingEmail(&emailPendingKey, &emailData)
    waitGroup.Done()
  }()
  go func () {
    authEmail, _ := utils.GetAuthEmailTemplate(&emailPendingKey)
    emailDeliveryData := types.EmailDelivery{}
    emailDeliveryData.Recipients = []string{emailData.Sender}
    emailDeliveryData.Body = authEmail
    utils.RequestEmailDelivery(emailDeliveryData)
    waitGroup.Done()
  }()
  waitGroup.Wait()
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "text/plain")
  w.Write([]byte(emailPendingKey))
}

// Handle Http response failure
func httpFailure(w http.ResponseWriter, httpStatusCode int) {
  http.Error(w, http.StatusText(httpStatusCode), httpStatusCode)
}
