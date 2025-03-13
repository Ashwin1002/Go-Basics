package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner   // or create anaynomous data
}

type owner struct {
	name string
	age  uint8
}

type electricEngine struct {
	kwh     uint8
	mpkwh uint8
	owner   // or create anaynomous data
}

type engine interface {
	milesLeft() uint8
}

func main() {
	//structs
	// looks similar to classes
	var motorEngine gasEngine = gasEngine{20, 15, owner{"hari", 18}}
	fmt.Println(motorEngine.mpg, motorEngine.gallons)

	fmt.Printf("Total miles left in tank: %v\n", motorEngine.milesLeft())
	remainingDistance(motorEngine, 50)

	var electricEngine electricEngine = electricEngine{100, 30, owner{"hari", 18}}
	fmt.Println(electricEngine.kwh, electricEngine.mpkwh)

	fmt.Printf("Total charge left in battery: %v\n", electricEngine.milesLeft())

	remainingDistance(electricEngine, 50)
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func remainingDistance(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can reach the destination!")
	} else {
		fmt.Println("You cannot reach the destination!")
	}
}
