package route

import "net/http"
import "encoding/json"
//import "github.com/donatien513/Contact-Me-Flow/utils"

type contactMeRequest struct {
  Sender string
  Message string
}

func AuthentificationHandler(w http.ResponseWriter, r *http.Request) {
  if r.Body == nil {
    httpFailure(w, http.StatusBadRequest)
    return
  }
  decoder := json.NewDecoder(r.Body)
  var reqData contactMeRequest
  decodeErr := decoder.Decode(&reqData)
  if decodeErr != nil {
    httpFailure(w, http.StatusBadRequest)
    return
  }

  if reqData.Sender == "" || reqData.Message == "" {
    httpFailure(w, http.StatusBadRequest)
    return
  }



  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "text/plain")
  w.Write([]byte("I am working :)"))
}

func httpFailure(w http.ResponseWriter, httpStatusCode int) {
  http.Error(w, http.StatusText(httpStatusCode), httpStatusCode)
}

