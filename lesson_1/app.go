package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	var a int8 // -128 <> 127, 1 byte, 8 bit
	var b int16 // -32768 <> 32767, 2 byte, 16 bit
	var c int32 // 4 byte, 32 bit
	var d int64 // 8 byte, 64 bit

	var e uint8 // 0 <> 255, 1 byte
	var f uint16 // 0 <> 65534, 2 byte
	var g uint32 // 4 byte
	var h uint64 // 8 byte

	var i byte // synonym uint8
	var j rune // synonym int32
	var k int // int32 or int64
	var l uint // uint32 or uint64

	var m float32 // 4 byte
	var n float64 // 8 byte

	var o complex64 = 1+2i
	var p complex128 = 4+90i

	var q bool

	var r string

	fmt.Println(a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r)

	var firstname = "FirstName"
	lastname := "LastName"
	var position string
	fmt.Println(firstname, lastname, "Default position:", position)
	position = "developer"
	fmt.Println("Position:", position)

	var (
		name = "TestName"
		age = 15
	)
	fmt.Println("Name:", name, "Age:", age)

	var country, countryID = "Estonia", 23
	str := fmt.Sprintf("The country is: %s, the country identifier is: %d", country, countryID)
	fmt.Println(str)
}