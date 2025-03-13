package main

import (
	"errors"
	"fmt"
)

func main() {

	printSomething("Hello World!")

	var result, remainder, err = division(4, 0)
	// using if-else conditional statement
	// if err != nil {
	// 	fmt.Println(err)
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of integer division is %v", result)
	// } else {
	// 	fmt.Printf("The result of integetr division is %v with remainder %v", result, remainder)
	// }

	switch {
	case err != nil:
		fmt.Println(err)
	case remainder == 0:
		fmt.Printf("The result of integer division is %v", result)
	default:
		fmt.Printf("The result of integer division is %v with remainder %v", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Printf("The division was exact")
	case 1, 2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}

}

func printSomething(value string) {
	fmt.Println(value)
}

func division(numerator int, denomitator int) (int, int, error) {
	var err error
	if denomitator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	quotient := numerator / denomitator
	remainder := numerator % denomitator
	return quotient, remainder, err
}
