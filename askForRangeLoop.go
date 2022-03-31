package main

import "fmt"

func main() {
	//Insert code for only even here
	var num1 int
	var num2 int
	fmt.Println("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&num2)

	for i := num1; i <= num2; i++ {
		if i%2 == 0 {
			fmt.Println("Even", i)
		}
		if i%2 != 0 {
			fmt.Println("Odd", i)
		}
	}

	if num1 > num2 {
		for i := num2; i <= num1; i++ {
			if i%2 == 0 {
				fmt.Println("Even", i)
			}
			if i%2 != 0 {
				fmt.Println("Odd", i)
			}
		}
	}
}
