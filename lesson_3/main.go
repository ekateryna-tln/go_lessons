package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p *Point) printPoint() {
	fmt.Println(p.X, p.Y)
}

func main() {
	pointers()
	structs()
	p1 := Point{
		X: 3,
		Y: 3,
	}
	p1.printPoint()
	p2 := &p1
	p2.printPoint()
}

func structs() {
	p1 := Point{
		X: 1,
		Y: 1,
	}
	fmt.Println(p1)

	p2 := Point{
		X: 2,
	}
	fmt.Println(p2)

	pp1 := &p1
	fmt.Println(pp1, pp1.X, pp1.Y)
	fmt.Println(*pp1)
	fmt.Println(*&pp1)
}

func pointers() {
	str := "hello world"
	num := 12

	fmt.Println(str, num)
	p := &str // get pointer

	fmt.Println(p, *p) // *p - get value

	*p = "new text"
	fmt.Println(*p, str)
}
