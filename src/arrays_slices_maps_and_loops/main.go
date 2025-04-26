package main

import "fmt"

func main() {
	// ARRAY

	// intArr := [3]int32{1,2,3}
	var intArr [3]int        // = [3]int{1, 2, 3} to set values in the positions
	fmt.Println(intArr[1:3]) // goes to position 1 until position 3

	var intSlice []int32 = []int32{1, 2, 3}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 4) // add the value 4 into array intSlice
	fmt.Printf("\nThe new length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	fmt.Println("------------------------------------------------------")

	// MAP

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])

	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Println("The age is ", age)
	} else {
		fmt.Println("The age is not exist")
	}

	delete(myMap2, "Jason") // delete(map, key)

	fmt.Println("------------------------------------------------------")

	// LOOP

	for name, age := range myMap2 {
		fmt.Printf("The name is %v and the age is %v\n", name, age)
	}

	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

}
