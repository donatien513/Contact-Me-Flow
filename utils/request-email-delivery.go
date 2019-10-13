package utils

import "bytes"
import "net/http"
import "encoding/json"
import "github.com/donatien513/Contact-Me-Flow/types"

var httpClient = &http.Client{}

func RequestEmailDelivery(emailDeliveryData types.EmailDelivery) error {
  jsonBytes := new(bytes.Buffer)
  jsonEncoder := json.NewEncoder(jsonBytes)
  jsonEncoder.Encode(emailDeliveryData)
  req, reqInitErr := http.NewRequest("POST", emailSenderURL, jsonBytes)
  if reqInitErr != nil {
    return reqInitErr
  }
  req.Header.Add("Authorization", emailSenderAuthToken)
  req.Header.Add("Content-Type", "application/json")
  resp, reqExecuteErr := httpClient.Do(req)
  print(resp.StatusCode)
  return reqExecuteErr
}
