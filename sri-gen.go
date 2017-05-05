package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "flag"
  "sync"
  "crypto/sha256"
  "crypto/sha512"
)

type sri struct {
  fileName  string
  hashType  int
  hash      []uint8
}

func main() {
  var wg sync.WaitGroup

  filePaths := os.Args[2:]
  hashType := flag.Int("hash", 256, "Hashing algorithm to use (256/384/512)");
  flag.Parse()

  results := make(chan sri, len(filePaths))
  hashFunction := hashChooser(hashType);

  for _, filePath := range filePaths {
    wg.Add(1)
    go func(filePath string) {
      defer wg.Done()
      hash := generateHash(filePath, hashFunction)
      resultData := sri{
        fileName: filePath,
        hash: hash,
        hashType: *hashType }
      results <- resultData
    } (filePath)
  }

  for i := 0; i < len(filePaths); i++ {
    fmt.Println(<-results)
  }

  wg.Wait()
}

func generateHash(path string, hashFunction func(data string) []uint8) []uint8 {
  fileContents := readFile(path)
  return hashFunction(fileContents)
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
