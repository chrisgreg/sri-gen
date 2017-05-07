package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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
	var sriResults = make([]sri, len(filePaths))
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
		sriResults[i] = <-results
	}

	writeToFile(sriResults)
}

func writeToFile(data []sri) {
	file, err := os.Create("result.txt")
	check(err)
	defer file.Close()
	fmt.Fprintf(file, "File\tHash\n\n")

	for _, result := range data {
		sriHash := "sha" + strconv.Itoa(result.hashType) + "-" + result.hash
		line := fmt.Sprintf("%s\t%s\n", result.fileName, sriHash)
		fmt.Fprintf(file, line)
	}

	fmt.Println("Finished creating file")
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
