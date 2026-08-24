package main

import (
	"fmt"
	"os"

	wc "github.com/ccwc/buildwc"
)

func main() {
	args := os.Args

	// check if there's a flag
	if len(args) < 1 {
		// if there's no flag, list out the help
		error := fmt.Errorf("command must contain flags: -w, -l or -e")
		fmt.Println(error)
		os.Exit(1)
	}

	// check if there's a file
	// if there's no file give a warning that there should be
	// and exit
	if len(args) > 2 && args[1] == "-c" {
		bytes, err := wc.CalculateBytes(args[2])
		if err != nil {
			fmt.Println(err)

		}

		fmt.Println(bytes, args[2])

	}

	if len(args) > 2 && args[1] == "-l" {
		lines, err := wc.CalculateLines(args[2])
		if err != nil {

			fmt.Println(err)
			os.Exit(2)
		}

		fmt.Println(lines, args[2])
	}

	if len(args) > 2 && args[1] == "-w" {
		words, err := wc.CalculateWords(args[2])
		if err != nil {

			fmt.Println(err)
			os.Exit(2)
		}

		fmt.Println(words, args[2])
	}

	if len(args) > 2 && args[1] == "-m" {
		runes, err := wc.CalculateRunes(args[2])
		if err != nil {

			fmt.Println(err)
			os.Exit(2)
		}

		fmt.Println(runes, args[2])
	}
}
