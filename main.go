package main

import (
	"fmt"
	"os"

	ccwc "github.com/ccwc/ccwc"
)

func main() {
	args := os.Args
	fmt.Println(args)

	// check if there's a flag
	// if there's no flag, list out the help

	// check if there's a file
	// if there's no file give a warning that there should be
	// and exit
	if len(args) > 2 && args[1] == "-c" {
		bytes, err := ccwc.CalculateBytes(args[2])
		if err != nil {
			fmt.Println(err)
			
		}

		fmt.Println(bytes, args[2])

	}
}