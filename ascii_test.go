package stl

import (
	"os"
	"testing"
)

func TestDecodeASCII(t *testing.T) {
	f, err := os.Open("testdata/sphericon.stl")
	if err != nil {
		t.Fatal(err)
	}
	_, err = DecodeASCII(f)
	if err != nil {
		t.Fatal(err)
	}
}
