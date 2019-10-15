package utils

import "net/http"

// Handle Http response failure
func HttpFailure(w http.ResponseWriter, httpStatusCode int) {
  http.Error(w, http.StatusText(httpStatusCode), httpStatusCode)
}
