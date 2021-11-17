package main

import "fmt"

const fromStr = "from"

func main() {
	// stack, last in first out
	defer fmt.Println("It is the last action")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	fmt.Println(helloWorld())
	fmt.Println(getCountry())

	h, w := helloWorld()
	c := getCountry()

	fmt.Println(h, w, fromStr, c)
	fmt.Println(getNumbers())

	fmt.Println("Full Cycle sum:", fullCycle())
	fmt.Println("Short Cycle sum:", shortCycle())
	fmt.Println("Analogue While sum:", analogueWhile())
	ifCondition()
	ifWithFunc()
	switchFunc()
	switchWithFunc()
	switchSimple()
}

func helloWorld() (string, string){
	a := "Hello"
	b := "World"
	return a, b
}

func getCountry() string {
	return "Estonia"
}

func getNumbers() (a, b, c int) {
	a, b, c = 1, 2, 3
	//this variant of using "return" is not recommended
	return
}

func fullCycle() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func shortCycle() int {
	sum := 0
	for ; sum < 100; {
		sum +=10
	}
	return sum
}

func analogueWhile() int {
	sum := 0
	for sum < 100 {
		sum += 10
	}
	return sum
}

func ifCondition() {
	a := 0
	for a < 3 {
		if a == 1 {
			fmt.Println("first step, a =", a)
		} else {
			fmt.Println("other step, a =", a)
		}
		a++
	}
}

func ifWithFunc() {
	if i := testFunc(); i == 1 {
		fmt.Println("if With Func")
	}
}

func testFunc() int {
	return 1
}

func switchFunc() {
	i := 0
	for i < 3 {
		switch i {
		case 1:
			fmt.Println("first step, i =", i)
		case 2:
			fmt.Println("second step, i =", i)
		default:
			fmt.Println("other step, i =", i)
		}
		i++
	}
}

func switchWithFunc() {
	switch i := testFuncForSwitch(1); i {
	case 1:
		fmt.Println("first step, i =", i)
	case 2:
		fmt.Println("second step, i =", i)
	default:
		fmt.Println("other step, i =", i)
	}
}

func switchSimple() {
	i := testFuncForSwitch(1);
	switch {
	case i == 1:
		fmt.Println("first step, i =", i)
	case i == 2:
		fmt.Println("second step, i =", i)
	default:
		fmt.Println("other step, i =", i)
	}
}

func testFuncForSwitch(i int) int {
	return i
}