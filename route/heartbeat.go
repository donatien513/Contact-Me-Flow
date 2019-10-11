package route

import "net/http"
import "io/ioutil"

func HeartbeatHandler(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "text/plain")
  resBody, _ := ioutil.ReadAll(r.Body)
  print(string(resBody))
  w.Write(resBody)
}
