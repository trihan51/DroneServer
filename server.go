package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

type Message struct {
     Command string "json:\"command\""
     Response string "json:\"response\""
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "path: %s. Welcome to the Drone's server.", r.URL.Path[:])
}

func connectHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("A device has contacted this drone server.")
  fmt.Println("Preparing to handle request.")

  p, err := ioutil.ReadAll(r.Body)
  if err != nil {
     panic(err)
  }

  var m Message
  err = json.Unmarshal(p, &m)

  if err != nil {
    panic(err)
  }

  fmt.Printf("Client request: %s", m.Command)

  responseMessage := Message{"No command", "Successful"}
  b, err := json.Marshal(responseMessage)

  if err != nil {
    panic(err)
  }

  w.Write(b)
}

func commandsHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "path: %s. Welcome to the Commands path.", r.URL.Path[:])
}

func main() {
  http.HandleFunc("/", defaultHandler)
  http.HandleFunc("/connect", connectHandler)
  http.HandleFunc("/commands", commandsHandler)
  http.ListenAndServe(":8080", nil)
}
