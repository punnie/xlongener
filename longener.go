package main

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
  url = "http://" + url // TODO(hpeixoto): do this right
  longened := Transform(url)

  go longener.kv.Store(longened, url)

  return longened
}

