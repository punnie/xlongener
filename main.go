package main

import "flag"

var filename = flag.String("filename", "zimbabwe.txt", "Path to the data store")
var interval = flag.Int("interval", 30, "Interval between data store disk writes")
var port     = flag.String("port", "4242", "Port for the webserver")

func main () {
  flag.Parse()
  // TODO(hpeixoto): Add an unix socket option
  LongenerHTTP(*filename, *interval, *port)
}

