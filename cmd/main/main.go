package main

import (
	"fmt"

	"github.com/Ashutoshbind15/projectstruct-go/internal/types"
	"github.com/Ashutoshbind15/projectstruct-go/internal/utils"
)

func main() {
	fmt.Println("base")

	c := utils.Adder(1,3)
	fmt.Println(c)

	box := types.Box {
		Size: 10,
	}

	fmt.Println(box)

	container := types.Container {
		Capacity: 5,
		Name: "containerid",
		Boxes: []types.Box{},
	}

	fmt.Println(container)

	utils.InsertNewBox(&container, 10)

	fmt.Println(container)
}