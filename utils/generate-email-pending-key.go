package utils

import "strconv"
import "math/rand"
import "time"

// Make strong random name
func GenerateEmailPendingKey() string {
  var result *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))
  var randomInt int = result.Int()
  return "email:" + strconv.Itoa(randomInt)
}