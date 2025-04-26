package main

import "fmt"
import "unicode/utf8"

func main() {
	var intNum int = 1
	fmt.Println("Using a int number:", intNum)

	var floatNum float64 = 1.5
	fmt.Println("Using a int number:", floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 * float32(intNum32)
	fmt.Println("Casting numeric types:", result)

	var myString string = "Hello World!"
	fmt.Println("Counting  the length:", utf8.RuneCountInString(myString))
	fmt.Println("Counting  the bytes:", len("°"))

	var myBool bool = true
	fmt.Println("Using bollean:", myBool)

	// Initializing variables
	myInt := 1
	myString2 := "Hello World!"
	myBool2 := true
	myFloat := 10.1
	fmt.Println(myInt, myString2, myBool2, myFloat)

	const myConst = "Hello World!"
	//myConst = "123" ERROR
}
