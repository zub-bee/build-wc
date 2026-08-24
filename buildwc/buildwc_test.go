package buildwc_test

import (
	"fmt"
	"os"
	"testing"

	wc "github.com/ccwc/buildwc"
)

func TestByte(t *testing.T) {
	args := os.Args
	got, _ := wc.CalculateBytes(args[2])
	want := "342190"

	if got != want {
		fmt.Printf("got %s wanted %s", got, want)
	}
}

func TestLine(t *testing.T) {

	args := os.Args
	got, _ := wc.CalculateLines(args[2])
	want := "7145"

	if got != want {
		fmt.Printf("got %s wanted %s", got, want)
	}
}
