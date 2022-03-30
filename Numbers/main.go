package main

import "fmt"

func main() {
	var number int
	fmt.Print("enter name")
	fmt.Scan(&number)
	for i:=0;i<=number;i++{
	if i%2==0{
	fmt.Print(i)
	}
}


