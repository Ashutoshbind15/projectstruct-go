package types

type Box struct {
	Size int
}

type Container struct {
	Capacity int
	Name     string
	Boxes    []Box
}