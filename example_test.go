package meshgrid_test

import (
	"fmt"

	"github.com/btracey/meshgrid"
	"github.com/gonum/floats"
)

func ExampleSingle() {
	fmt.Println("Create a grid using Span")
	grid := make([]float64, 3)
	floats.Span(grid, 0, 1)
	fmt.Println("grid =", grid)
	pts := meshgrid.Single(3, grid)
	fmt.Println("Generated points:")
	for _, v := range pts {
		fmt.Println(v)
	}
	// Output:
	// Create a grid using Span
	// grid = [0 0.5 1]
	// Generated points:
	// [0 0 0]
	// [0 0 0.5]
	// [0 0 1]
	// [0 0.5 0]
	// [0 0.5 0.5]
	// [0 0.5 1]
	// [0 1 0]
	// [0 1 0.5]
	// [0 1 1]
	// [0.5 0 0]
	// [0.5 0 0.5]
	// [0.5 0 1]
	// [0.5 0.5 0]
	// [0.5 0.5 0.5]
	// [0.5 0.5 1]
	// [0.5 1 0]
	// [0.5 1 0.5]
	// [0.5 1 1]
	// [1 0 0]
	// [1 0 0.5]
	// [1 0 1]
	// [1 0.5 0]
	// [1 0.5 0.5]
	// [1 0.5 1]
	// [1 1 0]
	// [1 1 0.5]
	// [1 1 1]
}

func ExampleMultiple() {
	grid := [][]float64{{0, 1}, {0, 0.5, 1}}
	pts := meshgrid.Multiple(grid)
	fmt.Println("Generated points:")
	for _, v := range pts {
		fmt.Println(v)
	}
	// Output:
	// Generated points:
	// [0 0]
	// [0 0.5]
	// [0 1]
	// [1 0]
	// [1 0.5]
	// [1 1]
}
