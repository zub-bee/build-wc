package main

import (
	"fmt"
	"os"
	"strings"

	wc "github.com/ccwc/buildwc"
)

const helpText = `
HELP:
	ccwc -[FLAG] file

flags
	-l calculate-lines
	-w calculate-words
	-c calculate-bytes
	-m calculate-runes`

func runCalculations(flag string, fileHandler *os.File, filePath string) {

	if flag == "-c" {

		bytes, err := wc.CalculateBytes(fileHandler)

		if err != nil {
			fmt.Println(err)

		}

		fmt.Println(bytes, filePath)

	}

	if flag == "-l" {

		lines, err := wc.CalculateLines(fileHandler)
		if err != nil {

			fmt.Println(helpText)
			fmt.Println(err)
			os.Exit(2)
		}

		fmt.Println(lines, filePath)
	}

	if flag == "-w" {

		words, err := wc.CalculateWords(fileHandler)
		if err != nil {

			fmt.Println(helpText)
			fmt.Println(err)
			os.Exit(2)
		}

		fmt.Println(words, filePath)
	}

	if flag == "-m" {

		runes, err := wc.CalculateRunes(fileHandler)
		if err != nil {

			fmt.Println(helpText)
			fmt.Println(err)
			os.Exit(2)
		}

		fmt.Println(runes, filePath)
	}
}

func main() {

	args := os.Args
	fileHandler := os.Stdin
	filePath := ""

	// check if there's a file
	// if there's no file give a warning that there should be
	// and exit

	if len(args) <= 2 {
		if len(args) == 2 {
			if strings.Contains(args[1], ".") {

				filePath = args[1]
				var err error
				fileHandler, err = wc.OpenFile(filePath)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				defer fileHandler.Close()

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
				return
			} else {

				if strings.Contains(args[1], "-") {
					runCalculations(args[1], fileHandler, "")
					return
				}

				fmt.Println("err: this is not a valid file or flag")
				fmt.Println(helpText)
				os.Exit(2)
			}
		}

		fmt.Println("err: no valid file present")
		fmt.Println(helpText)
		os.Exit(2)

	}

	filePath = args[len(args)-1]
	var err error
	fileHandler, err = wc.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fileHandler.Close()

	runCalculations(args[0], fileHandler, filePath)

}
