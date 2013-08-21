package main

import "testing"

func TestHashizeSize (t *testing.T) {
  hash := Hashize([]byte("hello"))
  if 256/8 != len(hash) {
    t.Errorf("Hash size wrong: Expected %d, received %d", 256/8, len(hash))
  }
}
