# go-stl

[![CI](https://github.com/neilpa/go-stl/workflows/CI/badge.svg)](https://github.com/neilpa/go-stl/actions)
[![GoDoc](https://godoc.org/neilpa.me/go-stl?status.svg)](https://godoc.org/neilpa.me/go-stl)

Simple go library for decoding binary and ASCII STL files as well as binary encoder.

## Usage

Reading triangle vertices of a binary or ASCII STL file

```go
package main

import (
    "fmt"
    "log"
    "os"

    "neilpa.me/go-stl"
)

func main() {
    f, err := os.Open("path/to/mesh.stl")
    if err != nil {
        log.Fatal(err)
    }

    mesh, err := stl.Decode(f)
    if err != nil {
        log.Fatal(err)
    }

    for _, face := range mesh.Faces {
        fmt.Println(face.Verts)
    }
}
```

See also [stlbox](cmd/stlbox/main.go) for an example of calculating
the bounding box.
