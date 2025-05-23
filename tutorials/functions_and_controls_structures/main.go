package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World!"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 0
	var result, remainder, err = intDivision(numerator, denominator)

	/* IF ELSE
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}*/

	// Switch without variable
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}

	// Switch with variable
	switch remainder {
	case 0:
		fmt.Printf("The division was exactly zero")
	case 1, 2: //1 or 2
		fmt.Printf("The division was one or two")
	default:
		fmt.Printf("The division was another number")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = errors.New("ERROR: Cannot divide by zero")
		return 0, 0, err
	}

	var result = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
