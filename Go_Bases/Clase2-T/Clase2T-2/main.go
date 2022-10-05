package main

import "fmt"

type matrix struct {
	Values    []float64
	Height    int
	Width     int
	Quadratic bool
	Maximum   float64
}

func (m *matrix) Set(values ...float64) {
	m.Values = values
}
func (m matrix) Get() {
	c := 0
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			fmt.Print(m.Values[c], "  ")
			c++
		}
		fmt.Println()
	}
}
func main() {
	ma := matrix{Height: 2, Width: 3, Quadratic: false, Maximum: 5.9}
	ma.Set(2.3, 4.1, 1.1, 4.7, 5.5, 5.9)
	ma.Get()
}
