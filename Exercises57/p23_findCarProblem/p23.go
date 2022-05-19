package p23

import (
	"fmt"
)

type Input struct {
}

type Output struct {
}

func RunP23() {

	//var input Input
	var s string

	fmt.Println("Is the car silent when you turn the key?")
	printInputRemind()
	fmt.Scanln(&s)

	carTroubleShooting(s)

	//output := Practice23Normal(input)

}

func Practice23Normal(input Input) (output Output) {

	return
}
