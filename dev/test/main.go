package main

import (
  "os"
  "log"
  "fmt"
  "net/http"
  "io/ioutil"
  "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load(fmt.Sprintf("./env/%s.env", os.Getenv("GO_ENV")))
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  fitbit_auth := os.Getenv("FITBIT_AUTH")
  fitbit_api_url_core := "https://api.fitbit.com/1/"

  client := &http.Client {}

  req, err := http.NewRequest("GET", fitbit_api_url_core + "user/-/profile.json", nil)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Authorization", "Bearer " + fitbit_auth)

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