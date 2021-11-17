package main

import "fmt"

func testCallback(callback func(int, int) int, s string) int {
	result := callback(5, 1)
	fmt.Println(s)
	return result
}

func main()  {
	sumCallback := func(n1, n2 int) int {
		return n1+n2
	}

	fmt.Println(testCallback(sumCallback, "plus"))

	multipleCallback := func(n1, n2 int) int {
		return n1*n2
	}

	fmt.Println(testCallback(multipleCallback, "multiple"))

}
