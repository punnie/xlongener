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
  if key == "" {
    http.ServeFile(handler.writer, handler.request, "index.html")
    return
  }

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

func LongenerHTTP (filename string, save_interval int, port string) {
  kv := KeyValue{}
  kv.Init(filename, time.Duration(save_interval) * time.Second)

  longener := Longener{&kv}

  http.HandleFunc("/create", generateHandler(longener, createHandler));
  http.HandleFunc("/", generateHandler(longener, fetchHandler));
  http.ListenAndServe(":" + port, nil)
}


