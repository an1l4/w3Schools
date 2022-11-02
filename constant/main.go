package main

import (
	"fmt"
	"math"
)

const PI = 3.14

const s string = "anila"

func main() {
	const ANIMAL = "cat"
	fmt.Printf("value of ANIMAL is %s\n", ANIMAL)
	const Correct = true
	fmt.Println("go rules?", Correct)

	fmt.Println("Happy", PI, "Day")

	const A = "GFG"
	var b = "GeekForGeeks"

	// Concat strings.
	var helloworld = A + "" + b
	helloworld += "!"
	fmt.Println(helloworld)

	// Compare strings.
	fmt.Println(A == "GFG")
	fmt.Println(b < A)

	const trueConst = true

	type myBool bool
	var defaultBool = trueConst
	var customBool myBool = trueConst

	fmt.Println(defaultBool)
	fmt.Println(customBool)

	fmt.Println(s)
	const n = 5
	const d = 395856 / n
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

}
