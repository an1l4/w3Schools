package main

import "fmt"

func main() {
	array := [6]int{44, 67, 23, 77, 89, 45}

	large := array[0]

	for i := 0; i <= 5; i++ {
		if large < array[i] {
			large = array[i]
		}
	}
	fmt.Println("largest number is ", large)

}
