package stl

import (
	"fmt"
	"os"
	"testing"
)

func TestDecodeASCII(t *testing.T) { // TODO Actual validiation and failure testing
	f, err := os.Open("testdata/sphericon.stl")
	if err != nil {
		t.Fatal(err)
	}
	file, err := DecodeASCII(f)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(file.Header)
}
