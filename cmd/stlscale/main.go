package main

import (
	"flag"
	"fmt"
	"os"

	"neilpa.me/go-stl"
)

var out = flag.String("o", "", "Output file for the scaled model")
var scale = flag.Float64("s", 1, "Scalar multiplier for the model")
// TODO Binary vs Ascii output

func main() {
	flag.Parse()
	if flag.NArg() == 0 || *out == "" {
		fmt.Fprintf(os.Stderr, "usage: %s -o OUTFILE [-s SCALE] INFILE\n", os.Args[0])
		os.Exit(2)
	}

	// TODO Support for stdin here as well?
	r, err := os.Open(flag.Arg(0))
	check(err)

	w, err := os.Create(*out)
	check(err)

	mesh, err := stl.Decode(r)
	check(err)

	encoder, err := stl.NewBinaryEncoder(w, "", len(mesh.Faces))
	check(err)

	for _, face := range mesh.Faces {
		face.Normal = [3]float32{0,0,0}
		for i := range face.Verts {
			for j := range face.Verts[i] {
				face.Verts[i][j] *= float32(*scale)
			}
		}
		err = encoder.WriteFace(face)
		check(err)
	}
}

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s", os.Args[0], err)
		os.Exit(1)
	}
}
