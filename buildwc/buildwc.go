package buildwc

import (
	"bufio"
	"fmt"
	"os"
)

func CalculateBytes(file string) (string, error) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	byteValue := fmt.Sprintf("%d", fileInfo.Size())
	return byteValue, nil
}

func CalculateLines(file string) (string, error) {
	fileHandle, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer fileHandle.Close()

	scanner := bufio.NewScanner(fileHandle)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", lineCount), nil
}

func CalculateWords(file string) (string, error) {
	fileHandle, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer fileHandle.Close()

	scanner := bufio.NewScanner(fileHandle)
	scanner.Split(bufio.ScanWords)
	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", wordCount), nil
}
