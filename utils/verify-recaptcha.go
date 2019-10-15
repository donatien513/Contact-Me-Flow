package utils

import "net/http"
import "github.com/tomasen/realip"
import "github.com/dpapathanasiou/go-recaptcha"
import "github.com/donatien513/Contact-Me-Flow/config"

func VerifyRecaptcha(r *http.Request) bool {
  var recaptchaResponse = r.Header.Get("x-recaptcha-response")
  var clientIP string = realip.FromRequest(r)
  if responseFound {
    result, err := recaptcha.Confirm(clientIP, recaptchaResponse)
    if err != nil {
      return false
    }
    return result
  }
  return false
}