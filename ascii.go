package stl

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
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

// scanState is used as a simple state machine for parsing the STL file
type scanState int
const (
	// start is the state at the beginning of scanning
	start scanState = iota
	// solid is the state after scanning the `solid` keyword
	solid
	// facet is the state after scanning the `facet` keyword
	facet
	// normal is the state after scanning the `normal` keyword
	normal
	// facetloop is the state after parsing the normal vector
	facetloop
	// outer is the state after scanning the `outer` keyword
	outer
	// loop is the state after scanning the `loop` keyword
	loop
	// vertex is the state after scanning the `vertex` keyword
	vertex
	// endloop is the state after scanning the `endloop` keyword
	endloop
	// endfacet is the state when looking for the next facet, e.g. after `endfacet` or `solid name`
	endfacet
)

// DecodeASCII parses all the faces from an STL text file.
func DecodeASCII(r io.Reader) (*File, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	file := File{}
	var face *Face
	var verts [][3]float32

	state := start
	for scanner.Scan() {
		word := scanner.Text()
		switch state {
		case start:
			if word != "solid" {
				return nil, fmt.Errorf("ASCII STL file must start with `solid`")
			}
			state = solid

		case solid:
			if word == "facet" {
				// The solid name is optional so explicitly check for facet
				state = facet
			} else {
				// Use the comment field for the name (TODO truncation)
				copy(file.Header.Comment[:], word)
				// endfacet begins the scan for the next face
				state = endfacet
			}

		case endfacet:
			switch word {
			case "facet":
				// Start a new face
				face = &Face{}
				verts = nil
				state = facet
			case "endsolid":
				// Completed parsing the mesh
				return &file, nil
			default:
				return nil, fmt.Errorf("Expected new `facet` or `endsolid`")
			}

		case facet:
			if word != "normal" {
				return nil, fmt.Errorf("Expected the start of a `normal`")
			}
			state = normal

		case normal:
			err := scanTriple(scanner, face.Normal[:])
			if err != nil {
				return nil, err
			}
			state = facetloop

		case facetloop:
			if word != "outer" {
				return nil, fmt.Errorf("Expected keywords `outer loop`")
			}
			state = outer

		case outer:
			if word != "loop" {
				return nil, fmt.Errorf("Expected keywords `outer loop`")
			}
			state = loop

		case loop:
			switch word {
			case "vertex":
				state = vertex
			case "endloop":
				state = endloop
			default:
				return nil, fmt.Errorf("Expected `vertex` or `endloop`")
			}

		case vertex:
			var v [3]float32
			err := scanTriple(scanner, v[:])
			if err != nil {
				return nil, err
			}
			verts = append(verts, v)
			// continue looping
			state = loop

		case endloop:
			if word != "endfacet" {
				return nil, fmt.Errorf("Expected keyword `endfacet`")
			}
			if len(verts) != 3 {
				return nil, fmt.Errorf("Expected 3 vertices")
			}
			copy(face.Verts[:], verts)
			file.Faces = append(file.Faces, *face)
			state = endfacet
		}
	}

	// Some stream reading failure no explicit `endsolid`
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return nil, fmt.Errorf("Unexpected EOF")
}

func scanTriple(scanner *bufio.Scanner, v []float32) error {
	// The scanner is already positioned at the first float
	x, err := strconv.ParseFloat(scanner.Text(), 32)
	if err != nil {
		return err
	}
	y, err := scanFloat32(scanner)
	if err != nil {
		return err
	}
	z, err := scanFloat32(scanner)
	if err != nil {
		return err
	}
	v[0], v[1], v[2] = float32(x), y, z
	return nil
}

func scanFloat32(scanner *bufio.Scanner) (float32, error) {
	if !scanner.Scan() {
		return 0, fmt.Errorf("Unexpected EOF")
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	n, err := strconv.ParseFloat(scanner.Text(), 32)
	return float32(n), err
}

