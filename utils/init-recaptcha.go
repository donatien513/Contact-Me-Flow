package utils

import "github.com/donatien513/Contact-Me-Flow/config"

func InitRecaptcha() {
  recaptcha.Init(config.RecaptchaSecretKey)
}