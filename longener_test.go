package main

import "testing"

func AssertEqual (a string, b string, t *testing.T) {
  if a != b {
    t.Errorf("Error: expected %s, got %s.", a, b)
  }
}

func TestSanitizer (t *testing.T) {
  AssertEqual("http://pokemon.com", Sanitize("pokemon.com"), t)
  AssertEqual("http://pokemon.com", Sanitize("pokemon.com "), t)
  AssertEqual("http://pokemon.com", Sanitize("pokemon.com\n"), t)
  AssertEqual("http://pokemon.com", Sanitize("http://pokemon.com"), t)
  AssertEqual("http://pokemon.com", Sanitize("  http://pokemon.com  "), t)
}

