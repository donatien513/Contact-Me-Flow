package utils

import "bytes"
import "io/ioutil"
import "net/http"
import "encoding/json"
import "github.com/donatien513/Contact-Me-Flow/types"

func GetAuthEmailTemplate(authEmailTemplateRequest *types.AuthEmailTemplateRequest) (string, error) {
  jsonBytes := new(bytes.Buffer)
  jsonEncoder := json.NewEncoder(jsonBytes)
  jsonEncoder.Encode(authEmailTemplateRequest)
  req, reqInitErr := http.NewRequest("POST", templateAuthURL, jsonBytes)
  if reqInitErr != nil {
    return "", reqInitErr
  }
  req.Header.Add("Authorization", emailSenderAuthToken)
  req.Header.Add("Content-Type", "application/json")
  resp, reqExecuteErr := httpClient.Do(req)
  if reqExecuteErr != nil {
    return "", reqExecuteErr
  }
  respBodyBytes, streamReadErr := ioutil.ReadAll(resp.Body)
  if streamReadErr != nil {
    return "", streamReadErr
  }
  bodyString := string(respBodyBytes)
  return bodyString, nil
}

