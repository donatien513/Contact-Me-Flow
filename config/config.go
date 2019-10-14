package config

import "os"

var ValidationHost string = os.Getenv("validation_host")
var TemplateAuthURL string = os.Getenv("template_auth_url")
var EmailSenderURL string = os.Getenv("email_sender_url")
var EmailSenderAuthToken string = os.Getenv("email_sender_auth_token")
var RedisAddr string = os.Getenv("redis_addr")
var RedisPass string = os.Getenv("redis_pass")
