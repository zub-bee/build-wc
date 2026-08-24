package buildwc

import (
	"fmt"
	"os"
)

func CalculateBytes(file string) (string, error) {
	FileInfo, err := os.Stat(file)

	if err != nil {
		return "", err
	}

	byteValue := fmt.Sprintf("%d", FileInfo.Size())
	return byteValue, nil
}