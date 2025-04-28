package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Println("The value of p is ", *p)
	fmt.Println("The value of i is ", i)

	p = &i // p -> i, store the address of i
	*p = 1 // the p value is 1, and changes the i variable

	fmt.Println("The value of p is ", *p)
	fmt.Println("The value of i is ", i)

}
