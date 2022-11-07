package main

import "fmt"

func main() {
	array1 := [...]int{1, 2, 3, 4, 5}
	array2 := [...]int{3, 5, 2, 5, 3}
	array3 := [...]int{1, 2, 3, 4, 5}

	if array1 == array2 {
		fmt.Println("array1 and array2 are same")
	} else {
		fmt.Println("array1 and array2 are not same")
	}

	if array1 == array3 {
		fmt.Println("array1 and array3 are same")
	} else {
		fmt.Println("array1 and array3 are not same")
	}

	if array2 == array3 {
		fmt.Println("array2 and array3 are same")
	} else {
		fmt.Println("array2 and array3 are not same")
	}

}
