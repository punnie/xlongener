package main

import "testing"
import "strings"


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
