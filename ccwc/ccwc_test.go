package buildwc_test

import (
	"fmt"
	"os"
	"testing"

	ccwc "github.com/ccwc/ccwc"
)

func TestByte(t *testing.T) {
	args := os.Args
	got, _ := ccwc.CalculateBytes(args[2])
	want := "342190"

	if got != want {
		fmt.Printf("got %s wanted %s", got, want)
	}
}