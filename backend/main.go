package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"regexp"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Data string `json:"hash" binding:"required"`
}

func main() {
	const Sha3Regex = `^([0-9a-fA-F]{128})$`
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.Use(CORSMiddleware())

	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {

		hash := c.Query("hash")
		if hash == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing hash parameter"})
			return
		}
		if ok, _ := regexp.MatchString(Sha3Regex, hash); !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hash format"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": processRequest(RequestBody{Data: hash})})

	})

	router.Run(":8080")
}

func processRequest(requestBody RequestBody) string {

	result := searchHash(requestBody.Data)
	if len(result) == 0 {
		return "No matches found"
	}

	return strings.Join(searchHash(requestBody.Data), ", ")
}

func searchHash(hash string) []string {
	const tablesPath = "./tables"
	files, err := os.ReadDir(tablesPath)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return nil
	}

	var wg sync.WaitGroup
	resultsChan := make(chan string, len(files))
	results := []string{}

	for _, file := range files {
		if !file.IsDir() {
			wg.Add(1)
			go func(file os.DirEntry) {
				defer wg.Done()
				filePath := filepath.Join(tablesPath, file.Name())
				f, err := os.Open(filePath)
				if err != nil {
					fmt.Printf("Error opening file %s: %v\n", file.Name(), err)
					return
				}
				defer f.Close()

				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					line := scanner.Text()
					if strings.Contains(line, hash) {
						resultsChan <- file.Name()
						break
					}
				}
				if err := scanner.Err(); err != nil {
					fmt.Printf("Error reading file %s: %v\n", file.Name(), err)
				}
			}(file)
		}
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	for fileName := range resultsChan {
		results = append(results, fileName)
	}

	return results
}
