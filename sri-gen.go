package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "flag"
  "crypto/sha512"
)

func main() {
  filePath := os.Args[1]
  




  fileContents := readFile(filePath)
  hash := create512Hash(fileContents)
  fmt.Printf("%x", hash)
}

func readFile(path string) string {
  content, err := ioutil.ReadFile(path)
  check(err)
  return string(content)
}

// func create256Hash(data string) []uint8 {
//   hash := sha256.New()
//   hash.Write([]byte(data))
//   return hash.Sum(nil)
// }

// func create384Hash(data string) []uint8 {
//   hash := sha512.New384()
//   hash.Write([]byte(data))
//   return hash.Sum(nil)
// }

func create512Hash(data string) []uint8 {
  hash := sha512.New()
  hash.Write([]byte(data))
  return hash.Sum(nil)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
