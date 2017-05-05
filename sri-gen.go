package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "flag"
  "crypto/sha256"
  "crypto/sha512"
)

func main() {
  filePath := os.Args[2]
  hashType := flag.Int("hash", 256, "Hashing algorithm to use (256/384/512)");
  flag.Parse()

  fileContents := readFile(filePath)
  hashFunction := hashChooser(hashType);

  hash := hashFunction(fileContents)

  fmt.Printf("%x", hash)
}

func hashChooser(hashType *int) func(data string) []uint8 {
  switch *hashType {
  case 256:
    return create256Hash
  case 384:
    return create384Hash
  case 512:
    return create512Hash
  }

  return create256Hash
}

func readFile(path string) string {
  content, err := ioutil.ReadFile(path)
  check(err)
  return string(content)
}

func create256Hash(data string) []uint8 {
  hash := sha256.New()
  hash.Write([]byte(data))
  return hash.Sum(nil)
}

func create384Hash(data string) []uint8 {
  hash := sha512.New384()
  hash.Write([]byte(data))
  return hash.Sum(nil)
}

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
