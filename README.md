# stl

Simple go library for decoding binary and ASCII STL files as well as binary encoder.

## Usage

Reading triangle vertices of a binary or ASCII STL file

```go
package main
import (
    "fmt"
    "log"
    "os"

    "neilpa.me/stl"
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

See also [stlbox](blob/master/cmd/stlbox/main.go) for an example of calculating
the bounding box.
