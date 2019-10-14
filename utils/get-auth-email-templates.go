package utils

import "bytes"
import "io/ioutil"
import "net/http"
import "net/url"
import "encoding/json"
import "github.com/donatien513/Contact-Me-Flow/config"

func GetAuthEmailTemplate(emailPendingKey *string) (string, error) {
  jsonBytes := new(bytes.Buffer)
  validationLink := makeValidationLink(emailPendingKey)
  postData := map[string]string{"validationLink": validationLink}
  jsonEncoder := json.NewEncoder(jsonBytes)
  jsonEncoder.Encode(postData)
  req, reqInitErr := http.NewRequest("POST", config.templateAuthURL, jsonBytes)
  if reqInitErr != nil {
    return "", reqInitErr
  }
  req.Header.Add("Authorization", config.emailSenderAuthToken)
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


func makeValidationLink(emailPendingKey *string) string {
  baseUrl, _ := url.Parse(config.validationHost)
  baseUrl.Path = "/send"
  params := url.Values{}
  params.Add("key", *emailPendingKey)
  baseUrl.RawQuery = params.Encode()
  return baseUrl.String()
}
