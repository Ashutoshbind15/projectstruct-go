package utils

import "github.com/Ashutoshbind15/projectstruct-go/internal/types"

func InsertNewBox(container *types.Container, sz int) {
	newBox := types.Box{
		Size: sz,
	}

	container.Boxes = append(container.Boxes, newBox)
}

func Adder(a, b int) int {
	return a + b
}