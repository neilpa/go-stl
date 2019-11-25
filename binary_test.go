package stl

import (
	"fmt"
	"os"
	"testing"
)

func TestDecodeBinary(t *testing.T) {
	f, err := os.Open("testdata/utah_teapot.stl")
	if err != nil {
		t.Fatal(err)
	}

	file, err := DecodeBinary(f)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(file.Header, len(file.Triangles))
}

func TestBinaryEncoder(t *testing.T) { // TODO
}
