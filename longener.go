package main

import "strings"

type KeyValuer interface {
  Store (key string, value string)
  Fetch (key string) string
}

type Longener struct {
  kv KeyValuer
}

func (longener *Longener) Fetch (key string) string {
  return longener.kv.Fetch(key)
}

func (longener *Longener) Store (url string) string {
  url = Sanitize(url)
  longened := Transform(url)

  go longener.kv.Store(longened, url)

  return longened
}

func Sanitize (url string) string {
  url = strings.TrimSpace(url)
  if strings.Contains(url, "://") {
    return url
  } else {
    return "http://" + url
  }
}

