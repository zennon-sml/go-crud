package main

import (
  "log"
  "database/sql"
  "fmt"
  "net/http"

  _ "github.com/lib/pq"
)

func handlerGet(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "get hello cow\n")
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "post hello cow\n")
}

func main() {
  const (
    host = "localhost"
    port = 5432
    user = "zennon"
    password = "asdf"
    dbname = "gocrud"
  )
  dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

  db, err := sql.Open("postgres", dsn)
  if err != nil {
    log.Fatal("failed to connect to db")
  } else {
    fmt.Println("success to connect")
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    log.Fatal("failed to ping the db")
  } else {
    fmt.Println("success to ping")
  }

  http.HandleFunc("GET /", handlerGet)
  http.HandleFunc("POST /", handlerPost)
  http.ListenAndServe(":8080", nil)
}
