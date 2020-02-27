package main

import "fmt"

func main() {
	for j := 0; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}
}
