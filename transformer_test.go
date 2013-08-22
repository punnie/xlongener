package main

import "testing"
import "bytes"

func TestHashizeSize (t *testing.T) {
  hash := Hashize([]byte("hello"))
  if 256/8 != len(hash) {
    t.Errorf("Hash size wrong: Expected %d, received %d", 256/8, len(hash))
  }
}

func TestHashize (t *testing.T) {
  if bytes.Equal(Hashize([]byte("bye")), Hashize([]byte("hello"))) {
    t.Errorf("Same hash")
  }
}

func TestTransform (t *testing.T) {
  if Transform("bye") == Transform("hello") {
    t.Errorf("Same hash")
  }
}

