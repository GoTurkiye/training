package main

type Calculator interface {
	Add(x, y int) int
	Subtract(x, y int) int
	Multiply(x, y int) int
	Divide(x, y int) float64
}

type Calculate struct{}

func (c Calculate) Add(x, y int) int {
	return x + y
}

func (c Calculate) Subtract(x, y int) int {
	return x - y
}

func (c Calculate) Multiply(x, y int) int {
	return x * y
}

func (c Calculate) Divide(x, y float64) float64 {
	return x/y
}
