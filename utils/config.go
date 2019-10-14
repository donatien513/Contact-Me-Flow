package utils

import "os"

var validationHost string = os.Getenv("validation_host")
var templateAuthURL string = os.Getenv("template_auth_url")
var emailSenderURL string = os.Getenv("email_sender_url")
var emailSenderAuthToken string = os.Getenv("email_sender_auth_token")
