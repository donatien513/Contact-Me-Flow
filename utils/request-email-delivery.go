package utils

import "io"
import "net/http"
import "encoding/json"
import "github.com/donatien513/Contact-Me-Flow/types"

var httpClient = &http.Client{}

func RequestEmailDelivery(emailData *types.EmailPendingRequest) error {
  emailDeliveryData := types.EmailDelivery{}
  emailDeliveryData.Recipients = []string{(*emailData).Sender}
  emailDeliveryData.Message = (*emailData).Message
  jsonOutput, jsonInput := io.Pipe()
  go func() {
    jsonEncoder := json.NewEncoder(jsonInput)
    defer jsonInput.Close()
    jsonEncoder.Encode(emailDeliveryData)
  }()
  req, reqInitErr := http.NewRequest("POST", emailSenderURL, jsonOutput)
  if reqInitErr != nil {
    return reqInitErr
  }
  req.Header.Add("authorization", emailSenderAuthToken)
  req.Header.Add("Content-Type", "application/json")
  resp, reqExecuteErr := httpClient.Do(req)
  print(resp.StatusCode)
  return reqExecuteErr
}
