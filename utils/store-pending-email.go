package utils

import "github.com/donatien513/Contact-Me-Flow/types"

var authWaitDuration time.Duration = time.Minute * 1

// Store pending email data to Redis
func StorePendingEmail(emailPendingKey *string, emailData *types.EmailPendingRequest) error {
  pipe := redisClient.TxPipeline()
  pipe.HSet(*emailPendingKey, "Sender", (*emailData).Sender);
  pipe.HSet(*emailPendingKey, "Message", (*emailData).Message);
  pipe.Expire(*emailPendingKey, authWaitDuration);
  _, err := pipe.Exec()
  return err
}
