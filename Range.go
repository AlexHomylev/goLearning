package main

import "fmt"

func main() {

	nums := []int{2, 3, 4, 15}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
}
