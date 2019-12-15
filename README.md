# go-stl

![Build Badge](https://github.com/neilpa/go-stl/workflows/CI/badge.svg)

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
