package stl

import (
	"fmt"
	"os"
	"testing"
)

func TestDecode(t *testing.T) { // TODO Actual validation
	tests := []struct {
		in string
	} {
		{ "sphericon.stl" },
		{ "utah_teapot.stl" },
	}
	for _, tt := range tests {
		t.Run(tt.in, func (t *testing.T) {
			f, err := os.Open("testdata/" + tt.in)
			if err != nil {
				t.Fatal(err)
			}
			file, err := Decode(f)
			if err != nil {
				t.Error(err)
			}
			fmt.Println(file.Header)
		})
	}
}
