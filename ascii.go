package stl

import (
	"fmt"
	"io"
)

// https://en.wikipedia.org/wiki/STL_%28file_format%29#ASCII_STL
//
// An ASCII STL file begins with the line
//
//		solid name
//
// where name is an optional string (though if name is omitted there must still be a space after
// solid). The file continues with any number of triangles, each represented as follows:
//
//		facet normal ni nj nk
//		    outer loop
//		        vertex v1x v1y v1z
//		        vertex v2x v2y v2z
//		        vertex v3x v3y v3z
//		    endloop
//		endfacet
//
// where each n or v is a floating-point number in sign-mantissa-"e"-sign-exponent format, e.g.,
// "2.648000e-002". The file concludes with
//
//		endsolid name
//
// The structure of the format suggests that other possibilities exist (e.g., facets with more than
// one "loop", or loops with more than three vertices). In practice, however, all facets are simple
// triangles.
//
// White space (spaces, tabs, newlines) may be used anywhere in the file except within numbers or
// words. The spaces between "facet" and "normal" and between "outer" and "loop" are required.

func DecodeASCII(r io.Reader) (*File, error) {
	return nil, fmt.Errorf("TODO - decode ASCII")
}
