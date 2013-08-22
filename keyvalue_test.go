package main

import "testing"
import "strings"
import "time"


func roundtrip (s string) bool { return s == Unstuff(Stuff(s)) }

func TestStuffing (t *testing.T) {

  cases := []string{"hello", "hello!", "hello#!", "hello, world!"}

  for _, test_case := range cases {
    if !roundtrip(test_case) {
      t.Error(test_case + " " + Stuff(test_case))
    }
    if strings.Contains(Stuff(test_case), ",") {
      t.Error(test_case + " " + Stuff(test_case))
    }
  }
}

func TestKV (t *testing.T) {
  kv := KeyValue{}
  kv.Init("bazonga", 10*time.Second)

  kv.Store("first", "primeiro")
  kv.Store("second", "segundo")

  if kv.Fetch("first") != "primeiro" {
      t.Error("first: " + kv.Fetch("first"))
  }

  if kv.Fetch("second") != "segundo" {
      t.Error("second: " + kv.Fetch("second"))
  }
}

