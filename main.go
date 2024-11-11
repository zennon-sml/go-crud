package main

import (
  "fmt"
  "net/http"
)

func handlerGet(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "get hello cow\n")
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "post hello cow\n")
}

func main() {
  http.HandleFunc("GET /", handlerGet)
  http.HandleFunc("POST /", handlerPost)
  http.ListenAndServe(":8080", nil)
}
