package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://mr-ppr79bybgkq.messaging.solace.cloud:9443/TechCody_YT_Topic"
  method := "POST"

  payload := strings.NewReader(`hello Solace`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Authorization", "Basic c29sYWNlLWNsb3VkLWNsaWVudDpvcXI2bGs4czFmOW1pb2hwZGV1NDBpYmZqbg==")
  req.Header.Add("Content-Type", "text/plain")
  req.Header.Add("Cookie", "TSID=6aefdb7abf39a9d7")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}