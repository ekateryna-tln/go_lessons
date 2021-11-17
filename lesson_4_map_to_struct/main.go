package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Point1 struct {
	X int
	Y int
}

type Point2 struct {
	X int `mapstructure:"xx"`
	Y int `mapstructure:"yy"`
}

func main() {
	pointMap1 := map[string]int{
		"x": 1,
		"y": 1,
	}
	point1 := Point1{}
	mapstructure.Decode(pointMap1, &point1)
	fmt.Println(point1)

	pointMap2 := map[string]int{
		"xx": 2,
		"yy": 2,
	}
	point2 := Point2{}
	mapstructure.Decode(pointMap2, &point2)
	fmt.Println(point2)

	for k, v := range pointMap1 {
		fmt.Println(k,v)
	}
}
