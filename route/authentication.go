package route

import "time"
import "net/http"
import "encoding/json"
import "github.com/donatien513/Contact-Me-Flow/utils"
import "github.com/donatien513/Contact-Me-Flow/types"

var authWaitDuration time.Duration = time.Minute * 1

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
  writePendingEmailErr := writePendingEmail(&emailPendingKey, &emailData)
  if writePendingEmailErr != nil {
    httpFailure(w, http.StatusInternalServerError)
    return
  }
  emailDeliveryError := utils.RequestEmailDelivery(&emailData)
  if emailDeliveryError != nil {
    httpFailure(w, http.StatusInternalServerError)
    return 
  }
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "text/plain")
  w.Write([]byte(emailPendingKey))
}

// Handle Http response failure
func httpFailure(w http.ResponseWriter, httpStatusCode int) {
  http.Error(w, http.StatusText(httpStatusCode), httpStatusCode)
}



// Store pending email data to Redis
func writePendingEmail(emailPendingKey *string, emailData *types.EmailPendingRequest) error {
  pipe := utils.RedisClient.TxPipeline()
  pipe.HSet(*emailPendingKey, "Sender", (*emailData).Sender);
  pipe.HSet(*emailPendingKey, "Message", (*emailData).Message);
  pipe.Expire(*emailPendingKey, authWaitDuration);
  _, err := pipe.Exec()
  return err
}
