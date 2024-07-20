package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/crypto/sha3"
)

func main() {
	const sourcesFile = "../sources/_sources.json"
	const tablesPath = "../tables"
	var sources []Source

	sourcesFileReader, err := os.Open(sourcesFile)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", sourcesFile, err)
		return
	}
	defer sourcesFileReader.Close()

	jsonParser := json.NewDecoder(sourcesFileReader)
	if err = jsonParser.Decode(&sources); err != nil {
		fmt.Printf("Error parsing file %s: %v\n", sourcesFile, err)
		return
	}

	//iterate over sources
	for _, source := range sources {
		realPath := filepath.Join("../sources/", source.Path)
		files, err := os.ReadDir(realPath)
		if err != nil {
			fmt.Printf("Error reading directory: %v\n", err)
			return
		}
		fmt.Printf("Hello! Let's build the rainbow tables from %s!\n", source.Repo)
		//iterate over files in source
		for _, file := range files {
			processFile(realPath, tablesPath, file.Name())
		}

	}

}

func processFile(basePath, outputPath, fileName string) {
	inputFilePath := filepath.Join(basePath, fileName)
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", inputFilePath, err)
		return
	}
	defer inputFile.Close()

	outputFilePath := filepath.Join(outputPath, fileName)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", outputFilePath, err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	fmt.Printf("Processing file: %s\n", fileName)

	scanner := bufio.NewScanner(inputFile)
	hasher := sha3.New512()
	for scanner.Scan() {
		hasher.Reset()
		hasher.Write([]byte(scanner.Text()))
		hash := hasher.Sum(nil)
		_, err := writer.WriteString(fmt.Sprintf("%x\n", hash))
		if err != nil {
			fmt.Printf("Error writing hash to file %s: %v\n", outputFilePath, err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning file %s: %v\n", inputFilePath, err)
	}
}

type Source struct {
	Repo string `json:"repo"`
	Path string `json:"path"`
}
