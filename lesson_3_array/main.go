package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"

	fmt.Println(a)

	numbers := [...]int{1, 2, 3}
	fmt.Println(numbers)
}
