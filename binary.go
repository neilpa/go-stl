package stl

import (
	"encoding/binary"
	"io"
)

type File struct {
	Header
	Triangles []Triangle
}

type Header struct {
	Header [80]byte
	NumTriangles uint32
}

func (h Header) String() string {
	return string(h.Header[:])
}

type Triangle struct {
	Normal [3]float32
	Verts [3][3]float32

	AttributeByteCount uint16
}

func DecodeBinary(r io.Reader) (*File, error) {
	var file File
	err := binary.Read(r, binary.LittleEndian, &file.Header)
	if err != nil {
		return nil, err
	}
	file.Triangles = make([]Triangle, file.NumTriangles)
	err = binary.Read(r, binary.LittleEndian, file.Triangles)
	if err != nil {
		return nil, err
	}
	return &file, nil
}
