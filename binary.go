package stl

import (
	"encoding/binary"
	"io"
)

type File struct {
	Header
	Faces []Face
}

type Header struct {
	Header [80]byte
	NumTriangles uint32
}

func (h Header) String() string {
	return string(h.Header[:])
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
