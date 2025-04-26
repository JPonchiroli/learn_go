package structs_and_interfaces

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type owner struct {
	name string
	age  uint8
}
