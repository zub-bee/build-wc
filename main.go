package main

import (
	"fmt"
	"os"

	wc "github.com/ccwc/buildwc"
)

func main() {

	args := os.Args
	lastIn := len(args) - 1
	fileHandler := os.Stdin
	filePath := args[lastIn]

	var err error
	fileHandler, err = wc.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fileHandler.Close()

	if err != nil {
		fmt.Println(err)
	}

	defer fileHandler.Close()

	// check if there's a flag
	if len(args) == 2 {
		// if there's no flag, give the default
		lines, err := wc.CalculateLines(fileHandler)
		fileHandler.Seek(0, 0)
		words, err := wc.CalculateWords(fileHandler)
		fileHandler.Seek(0, 0)
		bytes, err := wc.CalculateBytes(fileHandler)

		if err != nil {

			// print helper text
			fmt.Println(err)
			os.Exit(2)
		}

		fmt.Println(lines, words, bytes, filePath)
	}

	// check if there's a file
	// if there's no file give a warning that there should be
	// and exit

	if len(args) > 2 && args[1] == "-c" {

		bytes, err := wc.CalculateBytes(fileHandler)

		if err != nil {
			fmt.Println(err)

		}

		fmt.Println(bytes, args[2])

	}

	if len(args) > 2 && args[1] == "-l" {

		lines, err := wc.CalculateLines(fileHandler)
		if err != nil {

			fmt.Println(err)
			os.Exit(2)
		}

		fmt.Println(lines, args[2])
	}

	if len(args) > 2 && args[1] == "-w" {

		words, err := wc.CalculateWords(fileHandler)
		if err != nil {

			fmt.Println(err)
			os.Exit(2)
		}

		fmt.Println(words, args[2])
	}

	if len(args) > 2 && args[1] == "-m" {

		runes, err := wc.CalculateRunes(fileHandler)
		if err != nil {

			fmt.Println(err)
			os.Exit(2)
		}

		fmt.Println(runes, args[2])
	}
}
