package structs_and_interfaces

import (
	"fmt"
)

func main() {
	var myEngine gasEngine = gasEngine{25, 15, owner{"John Doe", 45}}
	//myEngine.gallons = 5
	canMakeIt(myEngine, 50)
}

func (e gasEngine) milesLeft() uint8 { // Importing the interface
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 { // Importing the interface
	return e.kwh * e.mpkwh
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("You can't make it!")
	}
}
