package main

import (
  "net/http"
)

func main () {
  http.HandleFunc("/", createImage)
  http.ListenAndServe(":8080", nil)
}