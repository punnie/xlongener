package main

import "time"
import "os"
import "fmt"
import "io/ioutil"
import "strings"

type Request func(map[string]string)
type Comm chan Request 
type KeyValue struct {
  urls map[string]string
  c Comm
}

func (kv *KeyValue) Init (filename string, interval time.Duration) {
  kv.urls = make(map[string]string)
  kv.c    = make(Comm)

  go kv.server()

  kv.load(filename)
  go kv.saver(filename, interval)
}

func (kv *KeyValue) Store (key string, value string) {
  kv.c <- func(urls map[string]string) { urls[key] = value }
}

func (kv *KeyValue) Fetch (key string) string {
  response := make(chan string)

  kv.c <- func(urls map[string]string) { response <- urls[key] }

  return <-response
}

func save (urls map[string]string, filename string) {
  fp, err := os.Create(filename)
  if err != nil {
    fmt.Println(err)
    return
  }

  defer fp.Close()

  for key, url := range urls {
    fp.Write([]byte(fmt.Sprintf("%s,%s\n", Stuff(key), Stuff(url))))
  }
}

func (kv *KeyValue) load (filename string) {
  content, err := ioutil.ReadFile(filename)
  if err != nil {
    panic(err)
  }

  for _, entry := range strings.Split(string(content), "\n") {
    var things = strings.Split(entry, ",")
    kv.Store(Unstuff(things[0]), Unstuff(things[1]))
  }
}

func (kv *KeyValue) saver (filename string, interval time.Duration) {
  for {
    time.Sleep(interval)
    kv.c <- func(urls map[string]string) { save(urls, filename) }
  }
}

func (kv *KeyValue) server () {
  for {
    (<-kv.c)(kv.urls)
  }
}

func Stuff (s string) string {
  return strings.Replace(strings.Replace(s, "!", "#!", -1), ",", "!!", -1)
}

func Unstuff (s string) string {
  return strings.Replace(strings.Replace(s, "!!", ",", -1), "#!", "!", -1)
}

