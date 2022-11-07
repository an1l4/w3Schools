package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3, 4, 5}
	sum := 0

	for i := 0; i <= 4; i++ {
		sum = sum + array[i]
	}
	fmt.Println("sum is", sum)

}
