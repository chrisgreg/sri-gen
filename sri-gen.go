package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type sri struct {
	fileName string
	hashType int
	hash     string
}

func main() {

	filePaths := os.Args[2:]
	hashType := flag.Int("hash", 256, "Hashing algorithm to use (256/384/512)")
	flag.Parse()

	results := make(chan sri, len(filePaths))
	hashFunction := hashChooser(hashType)

	for _, filePath := range filePaths {
		go func(filePath string) {
			hash := generateHash(filePath, hashFunction)
			resultData := sri{
				fileName: filePath,
				hash:     hex.EncodeToString(hash),
				hashType: *hashType}
			results <- resultData
		}(filePath)
	}

	for i := 0; i < len(filePaths); i++ {
		fmt.Println(<-results)
	}
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
