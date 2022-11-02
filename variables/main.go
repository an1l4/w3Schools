package main

import "fmt"

func main() {
	// Variable declared and
	// initialized without the
	// explicit type
	var myVariable1 = 34
	fmt.Printf("value of myVariable1 is %d and type is %T\n", myVariable1, myVariable1)
	var myVariable2 = "anila"
	fmt.Printf("value of myVariable2 is %s and type is %T\n", myVariable2, myVariable2)
	var myVariable3 = 34.98
	fmt.Printf("value of myVariable3 is %f and typeis %T\n", myVariable3, myVariable3)

	// Variable declared and
	// initialized without expression
	var myVariable4 int
	fmt.Printf("value of myVariable4 is %d\n", myVariable4)
	var myVariable5 string
	fmt.Printf("value of myVariable5 is %s\n", myVariable5)
	var myVariable6 float64
	fmt.Printf("value of myVariable6 is %f\n", myVariable6)

	// Multiple variables of the same type
	// are declared and initialized
	// in the single line

	var myVariable7, myVariable8, myVariable9 int = 34, 67, 23
	fmt.Printf("value of myVariable7 is %d\n", myVariable7)
	fmt.Printf("value of myVariable8 is %d\n", myVariable8)
	fmt.Printf("value of myVariable is %d\n", myVariable9)

	// Multiple variables of different types
	// are declared and initialized in the single line
	var myVariable10, myVariable11, myVariable12 = 98, "geek", 56.95
	fmt.Printf("value of myVariable10 is %d and type is %T\n", myVariable10, myVariable10)
	fmt.Printf("value of myVariable11 is %s and type is %T\n", myVariable11, myVariable11)
	fmt.Printf("value of myVariable12 is %f and type is %T\n", myVariable12, myVariable12)

	// Using short variable declaration
	myVar1 := 23
	myVar2 := "ball"
	myVar3 := 45.87
	fmt.Printf("Value of myVar1 is %d and type is %T\n", myVar1, myVar1)
	fmt.Printf("value of myVar2 is %s and type is %T\n", myVar2, myVar2)
	fmt.Printf("value of myVar3 is %f and type is %T\n", myVar3, myVar3)

	// Using short variable declaration
	// Multiple variables of same types
	// are declared and initialized in
	// the single line
	myVar4, myVar5, myVar6 := 12, 56, 38
	fmt.Printf("value of myVar4 is %d and type is %T\n", myVar4, myVar4)
	fmt.Printf("value of myVar5 is %d and type is %T\n", myVar5, myVar5)
	fmt.Printf("value of myVar6 is %d amd type is %T\n", myVar6, myVar6)

	// Using short variable declaration
	// Here, short variable declaration acts
	// as an assignment for myvar8 variable
	// because same variable present in the same block
	// so the value of myvar8 is changed from 45 to 100
	myVar7, myVar8 := 39, 45
	myVar9, myVar8 := 45, 100
	fmt.Printf("value od myVar7 is %d and myVar8 is %d\n", myVar7, myVar8)
	fmt.Printf("value of myVar9 is %d and myVar8 is %d\n", myVar9, myVar8)
	// If you try to run the commented lines,
	// then compiler will gives error because
	// these variables are already defined
	// myvar1, myvar2 := 43, 47
	// myvar2:= 200

	// Using short variable declaration
	// Multiple variables of different types
	// are declared and initialized in the single line
	myVar10, myVar11, myVar12 := 23, "engineer", 10.85
	fmt.Printf("value of myVar10 is %d and type is %T\n", myVar10, myVar10)
	fmt.Printf("value of myVar11 is %s and type is %T\n", myVar11, myVar11)
	fmt.Printf("value of myVar12 is %f and type is %T\n", myVar12, myVar12)

}
