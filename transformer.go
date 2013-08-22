package main

import "crypto/sha256"

func Xize (x []byte) string {
  k := ""
  for i := uint(0); i < uint(8*len(x)); i++ {
    if (x[i/8] >> (i%8)) % 2 == 0 {
      k += "x"
    } else {
      k += "X"
    }
  }

  return k
}

func Hashize (x []byte) []byte {
  hasher := sha256.New()
  hasher.Write(x)
  return hasher.Sum(nil)
}

func Transform (message string) string {
  return Xize(Hashize([]byte(message)))
}

