package stl

import (
	"bytes"
	"io"
)

// Decode guesses whether the reader is a binary or ASCII STL file and
// attempts to decode it as such.
func Decode(r io.Reader) (*File, error) {
	buf := make([]byte, 5)
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}

	// Reset the stream for actual decoding
	if s, ok := r.(io.Seeker); ok {
		_, err = s.Seek(0, io.SeekStart)
		if err != nil {
			return nil, err
		}
	} else {
		// Don't skip the already read bytes
		r = io.MultiReader(bytes.NewReader(buf), r)
	}

	if string(buf) == "solid" {
		return DecodeASCII(r)
	} else {
		return DecodeBinary(r)
	}
}
