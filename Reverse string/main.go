package main

import "fmt"

func main() {
	var input string
	var result string
	fmt.Print("enter name\n")
	fmt.Scan(&input)
	for _, a := range input {
		result = string(a) + result
		fmt.Print(result, "\n")
	}

}
