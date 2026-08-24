package buildwc_test

import (
	"fmt"
	"os"
	"testing"

	wc "github.com/ccwc/buildwc"
)

var args = os.Args

const fileName = "../test.txt"

func TestByte(t *testing.T) {

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

func TestWords(t *testing.T) {

	want := "58164"
	got, err := wc.CalculateWords(fileName)

	if err != nil {
		t.Errorf("error running the words %q", err)
	}

	if got != want {
		t.Errorf("got %q words wanted %s words", got, want)
	}
}
