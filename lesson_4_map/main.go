package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p Point) printPoint() {
	fmt.Println(p.X, p.Y)
}

func main() {
	// variable declaration
	var pointMap map[string]Point

	// initialization
	pointMap = map[string]Point{}

	// declaration and initialization
	pointMap1 := map[string]Point{}
	intMap1 := map[string]int{}
	pointMap2 := make(map[string]Point)
	intMap2 := make(map[string]int)
	pointMapWithValue := map[int]Point{
		1: {
			X: 2,
			Y: 2,
		},
	}

	key := "first"
	pointMap1[key] = Point{
		X: 1,
		Y: 1,
	}
	pointMap1[key].printPoint()

	fmt.Println(pointMap, pointMap1, pointMap2, intMap1, intMap2, pointMapWithValue)

	// check if key exist
	value, ok := pointMap1[key]
	if ok {
		fmt.Println("Value exit:", value)
	}
}
