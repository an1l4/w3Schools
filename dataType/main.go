package main

import "fmt"

func main() {
	//uint - Depends on platform:32 bits in 32 bit systems and 64 bit in 64 bit systems
	//uint8 - 0 to 255
	//uint16 0 to 65535
	//uint32
	//uint64
	var x uint8 = 225
	fmt.Println(x, x-3)

	//int
	//int8 - -128 ti 127
	//int16 - -32768 to 32767
	//int32
	//int64
	var y int16 = 32767
	fmt.Println(y, y+1)

	//float32
	//float64
	a := 34.97
	b := 21.54
	fmt.Println(a - b)

	var c complex64 = complex(6, 2)
	var d complex128 = complex(9, 2)
	fmt.Println(c, d)

	str1 := "anila"
	str2 := "anu"
	str3 := "anila"

	result1 := str1 == str2
	result2 := str1 == str3
	result3 := str2 == str3

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	str := "geeks"
	length := len(str)
	fmt.Println(length)
	fmt.Println(str)
}
