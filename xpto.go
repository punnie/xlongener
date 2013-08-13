package main

import (
  "fmt"
  "net/http"
  "time"
)

type LongenerHandler struct {
  writer http.ResponseWriter
  request *http.Request
  longener Longener
}

type LongenerHandlerFunc func(LongenerHandler)

func createHandler (handler LongenerHandler) {
  value := handler.request.FormValue("url")
  key := handler.longener.Store(value)
  fmt.Fprintf(handler.writer, key)
}

func fetchHandler (handler LongenerHandler) {
  key := handler.request.URL.Path[1:]
  location := handler.longener.Fetch(key)

  if location == "" {
    fmt.Fprintf(handler.writer, "not found")
  } else {
    http.Redirect(handler.writer, handler.request, location, http.StatusFound)
  }
}

func generateHandler (longener Longener, handler LongenerHandlerFunc)  http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request) {
    handler(LongenerHandler{w, r, longener})
  }
}

func main () {
  // TODO(hpeixoto): grab this from program arguments
  filename      := "zimbabwe.txt"
  save_interval := 120 * time.Second
  port          := "4242"

  kv := KeyValue{}
  kv.Init(filename, save_interval)

  longener := Longener{&kv}

  http.HandleFunc("/create", generateHandler(longener, createHandler));
  http.HandleFunc("/", generateHandler(longener, fetchHandler));
  http.ListenAndServe(":" + port, nil)
}

