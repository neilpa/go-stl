package stl

import (
	"encoding/binary"
	"io"
)

type File struct {
	Header
	Faces []Face
}

const (
	commentSize = 80
	headerSize = commentSize+4
)

type Header struct {
	Comment [commentSize]byte
	NumTriangles uint32
}

func (h Header) String() string {
	return string(h.Comment[:])
}

type Face struct {
	Normal [3]float32
	// TODO Triangle Type
	Verts [3][3]float32

	AttributeByteCount uint16
}

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

type BinaryEncoder struct {
	ws io.WriteSeeker
	faces uint32
}

// NewBinaryEncoder creates an encoder that wraps the provided writer. It's
// important to call Close when done writing to ensure the face count is
// written the header.
func NewBinaryEncoder(ws io.WriteSeeker) (*BinaryEncoder, error) {
	_, err := ws.Seek(headerSize, io.SeekStart)
	if err != nil {
		return nil, err
	}
	return &BinaryEncoder{ws, 0}, nil
}

// WriteFace writes the face to the wrapped file.
func (e *BinaryEncoder) WriteFace(f Face) error {
	err := binary.Write(e.ws, binary.LittleEndian, f)
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

// Close writes the total face count. If the writer is also an io.Closer this
// will close the underlying stream.
func (e *BinaryEncoder) Close() error {
	_, err := e.ws.Seek(commentSize, io.SeekStart)
	if err != nil {
		return err
	}
	err = binary.Write(e.ws, binary.LittleEndian, e.faces)
	if err != nil {
		return err
	}
	if c, ok := e.ws.(io.Closer); ok {
		return c.Close()
	}
	return nil
}
