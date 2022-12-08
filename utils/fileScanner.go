package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetFileContent(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Scanner Error:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fileContent := []string{}
	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}
	return fileContent
}
