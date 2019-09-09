package web

import (
 "fmt"
 "net/http"
 "encoding/json"
)

// https://gobyexample.com/json
type Event struct {
  EvenType string `json:"eventType"`
}

func (e Event) toString() string {
  res, err := json.Marshal(e)

  if err != nil {
    fmt.Println(err)
  }
  return string(res)
}

func index_handler(w http.ResponseWriter, r *http.Request) {
  e := &Event{"test"}
  w.Header().Set("Content-Type", "application/json")
  fmt.Fprintf(w,e.toString())
}

func Run() {
  fmt.Println("=== go web ===")

  http.HandleFunc("/", index_handler)
  http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Whoa test")
  })
  http.ListenAndServe(":8000", nil)
}
