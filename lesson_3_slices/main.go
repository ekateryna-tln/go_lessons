package main

import "fmt"

//slice - это обертка над массивом
func main() {
	letters := []string{"a", "b", "c"}
	fmt.Println(letters)
	letters = append(letters, "d")
	letters = append(letters, "e")
	fmt.Println(letters)
	letters = append(letters, "h", "i", "j")
	fmt.Println(letters)

	words := make([]string, 3)
	words[0] = "name"
	words[1] = "age"
	words[2] = "birthday"
	fmt.Println(len(words), cap(words))
	// words[3] = "birthday" - panic
	words = append(words, "address")
	fmt.Println(len(words), cap(words))

	animalsArr := [4]string{
		"dog",
		"cat",
		"giraffe",
		"elephant",
	}

	// from 0 to 1
	a := animalsArr[0:2]
	fmt.Println(a)

	// from 1 to 2
	b := animalsArr[1:3]
	fmt.Println(b)

	animalsArr[0] = "other_dog"
	fmt.Println(a)

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := s[5:6]
	fmt.Println(t)

	start1 := s[0:5]
	start2:= s[:5]
	fmt.Println(start1, start2)

	end1 := s[5:10]
	end2 := s[5:]
	fmt.Println(end1, end2)

	all := s[:]
	fmt.Println(all)

	var sliceNil []int
	fmt.Println(sliceNil, len(sliceNil), cap(sliceNil))
	if sliceNil == nil {
		fmt.Println("slice is nil")
	}
}
