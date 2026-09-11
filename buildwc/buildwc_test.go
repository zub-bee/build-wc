package buildwc_test

import (
	"testing"

	wc "github.com/ccwc/buildwc"
)

func TestByte(t *testing.T) {

	var fileHandle, _ = wc.OpenFile("../test.txt")
	got, _ := wc.CalculateBytes(fileHandle)
	want := "342190"

	if got != want {
		t.Errorf("got %s wanted %s", got, want)
	}
}

func TestLine(t *testing.T) {

	var fileHandle, _ = wc.OpenFile("../test.txt")
	got, _ := wc.CalculateLines(fileHandle)
	want := "7145"

	if got != want {
		t.Errorf("got %s wanted %s", got, want)
	}
}

func TestWords(t *testing.T) {

	var fileHandle, _ = wc.OpenFile("../test.txt")
	want := "58164"
	got, err := wc.CalculateWords(fileHandle)

	if err != nil {
		t.Errorf("error running the words %q", err)
	}

	if got != want {
		t.Errorf("got %q words wanted %s words", got, want)
	}
}

func TestRunes(t *testing.T) {

	var fileHandle, _ = wc.OpenFile("../test.txt")
	want := "339292"
	got, err := wc.CalculateRunes(fileHandle)

	if err != nil {
		t.Errorf("error running the words %q", err)
	}

	if got != want {
		t.Errorf("got %q words wanted %s runes", got, want)
	}
}
