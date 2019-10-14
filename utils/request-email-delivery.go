package utils

import "bytes"
import "net/http"
import "encoding/json"
import "github.com/donatien513/Contact-Me-Flow/types"
import "github.com/donatien513/Contact-Me-Flow/config"

var httpClient = &http.Client{}

func RequestEmailDelivery(emailDeliveryData types.EmailDelivery) error {
  jsonBytes := new(bytes.Buffer)
  jsonEncoder := json.NewEncoder(jsonBytes)
  jsonEncoder.Encode(config.EmailDeliveryData)
  req, reqInitErr := http.NewRequest("POST", config.EmailSenderURL, jsonBytes)
  if reqInitErr != nil {
    return reqInitErr
  }
  req.Header.Add("Authorization", config.EmailSenderAuthToken)
  req.Header.Add("Content-Type", "application/json")
  resp, reqExecuteErr := httpClient.Do(req)
  print(resp.StatusCode)
  return reqExecuteErr
}
