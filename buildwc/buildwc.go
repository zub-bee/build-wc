package buildwc

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func OpenFile(file string) (*os.File, error) {

	fileHandle, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	return fileHandle, nil
}

func CalculateBytes(fileHandle io.Reader) (string, error) {

	scanner := bufio.NewScanner(fileHandle)
	scanner.Split(bufio.ScanBytes)
	byteCount := 0
	for scanner.Scan() {
		byteCount++
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", byteCount), nil
}

func CalculateLines(fileHandle io.Reader) (string, error) {

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

func CalculateWords(fileHandle io.Reader) (string, error) {

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

func CalculateRunes(fileHandle io.Reader) (string, error) {

	scanner := bufio.NewScanner(fileHandle)
	scanner.Split(bufio.ScanRunes)
	runeCount := 0
	for scanner.Scan() {
		runeCount++
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", runeCount), nil
}
