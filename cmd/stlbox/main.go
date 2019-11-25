package main

import (
	"fmt"
	"math"
	"os"

	"neilpa.me/stl"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s FILE\n", os.Args[0])
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s", os.Args[0], err)
		os.Exit(1)
	}
	mesh, err := stl.Decode(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s", os.Args[0], err)
		os.Exit(1)
	}

	var (
		minX float32 = math.MaxFloat32
		minY float32 = math.MaxFloat32
		minZ float32 = math.MaxFloat32
		maxX float32 = -math.MaxFloat32
		maxY float32 = -math.MaxFloat32
		maxZ float32 = -math.MaxFloat32
	)
	for _, face := range mesh.Faces {
		v := face.Verts
		minX = min(min(minX, v[0][0]), min(v[1][0], v[2][0]))
		minY = min(min(minY, v[0][1]), min(v[1][1], v[2][1]))
		minZ = min(min(minZ, v[0][2]), min(v[1][2], v[2][2]))
		maxX = max(max(maxX, v[0][0]), max(v[1][0], v[2][0]))
		maxY = max(max(maxY, v[0][1]), max(v[1][1], v[2][1]))
		maxZ = max(max(maxZ, v[0][2]), max(v[1][2], v[2][2]))
	}
	fmt.Printf("min-corner (%f,%f,%f)\n", minX, minY, minZ)
	fmt.Printf("max-corner (%f,%f,%f)\n", maxX, maxY, maxZ)
}

func min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}
