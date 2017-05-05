package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "crypto/sha256"
)

func main() {
  filePath := os.Args[1]
  fileContents := readFile(filePath)
  create256Hash(fileContents)
}

func readFile(path string) string {
  content, err := ioutil.ReadFile(path)
  check(err)
  return string(content)
}

func create256Hash(data string) {
  hash := sha256.New()
  hash.Write([]byte(data))
  fmt.Printf("%x", hash.Sum(nil))
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
