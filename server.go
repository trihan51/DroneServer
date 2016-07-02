package main

import (
  "fmt"
  "net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "path: %s. Welcome to the Drone's server.", r.URL.Path[:])
}

func connectHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "path: %s. Welcome to the Connect path.", r.URL.Path[:])
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
