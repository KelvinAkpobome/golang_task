package main

import "fmt"

func main() {
	fmt.Println("Hi, My name is Dectobin")
	fmt.Println("Please input Your Number")

	var userInput int
	fmt.Scan(&userInput)
	fmt.Printf("The binary format is %b\n", userInput)

}
