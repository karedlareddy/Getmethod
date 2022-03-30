package main

import "fmt"

func main() {
	var number int
	fmt.Println("enter number\n")

	fmt.Scan(&number)

	if i := number; i%2 == 0 {
		fmt.Println("it is an even number")
	} else {
		fmt.Println("it is an odd number")
	}
}
