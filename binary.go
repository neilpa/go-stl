package stl // import "neilpa.me/go-stl"

import (
	"encoding/binary"
	"fmt"
	"io"
)

// File represents a parsed STL file.
type File struct {
	Header
	Faces []Face
}

const (
	commentSize = 80
)

// Header is the metadata of a parsed STL file.
type Header struct {
	Comment [commentSize]byte
	NumTriangles uint32
}

func (h Header) String() string {
	return string(h.Comment[:])
}

// Face contains the vertex and normal data for a triangle in the STL mesh.
type Face struct {
	Normal [3]float32
	// TODO Triangle Type
	Verts [3][3]float32

	AttributeByteCount uint16
}

// DecodeBinary parses all the faces from an STL binary file.
func DecodeBinary(r io.Reader) (*File, error) {
	var file File
	err := binary.Read(r, binary.LittleEndian, &file.Header)
	if err != nil {
		return nil, err
	}
	file.Faces = make([]Face, file.NumTriangles)
	err = binary.Read(r, binary.LittleEndian, file.Faces)
	if err != nil {
		return nil, err
	}
	return &file, nil
}

// BinaryEncoder is used for serializing STL mesh data in the binary format.
type BinaryEncoder struct {
	w io.Writer
	s io.Seeker
	faces uint32
}

// NewBinaryEncoder creates an encoder that wraps the provided writer. Use -1
// if the number of faces are uknown at the start. In this case it's required
// the io.Writer also support io.Seeker so that they can be written at the
// end during Close().
func NewBinaryEncoder(w io.Writer, comment string, faces int) (*BinaryEncoder, error) {
	var header Header
	var seeker io.Seeker

	if faces < 0 {
		if s, ok := w.(io.Seeker); ok {
			seeker = s
		} else {
			return nil, fmt.Errorf("Must specify num faces or provide an io.Seeker")
		}
	} else {
		header.NumTriangles = uint32(faces)
	}

	copy(header.Comment[:], []byte(comment))
	err := binary.Write(w, binary.LittleEndian, header)
	if err != nil {
		return nil, err
	}

	return &BinaryEncoder{w, seeker, 0}, nil
}

// WriteFace writes the face to the wrapped file.
func (e *BinaryEncoder) WriteFace(f Face) error {
	err := binary.Write(e.w, binary.LittleEndian, f)
	if err != nil {
		return err
	}
	e.faces++
	return nil
}

// WriteTriangle writes a new face for the given triangle points wihtout
// calculating the normal.
func (e *BinaryEncoder) WriteTriangle(a, b, c [3]float32) error {
	return e.WriteFace(Face{Verts: [3][3]float32{ a, b, c }})
}

// Close writes the total face count if it wasn't provided up front. If
// the writer is also an io.Closer this will close the underlying stream.
func (e *BinaryEncoder) Close() error {
	if e.s != nil {
		_, err := e.s.Seek(commentSize, io.SeekStart)
		if err != nil {
			return err
		}
		err = binary.Write(e.w, binary.LittleEndian, e.faces)
		if err != nil {
			return err
		}
	}
	if c, ok := e.w.(io.Closer); ok {
		return c.Close()
	}
	return nil
}
