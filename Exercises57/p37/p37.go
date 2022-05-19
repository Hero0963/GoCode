package p37

import (
	"fmt"
)

type Input struct {
	Len              int
	SpecialCharCount int
	NumberCount      int
}

type Output struct {
	Password string
}

func RunP37() {

	var input Input

	fmt.Println("What's the length?")
	fmt.Scanln(&input.Len)

	fmt.Println("How many special characters?")
	fmt.Scanln(&input.SpecialCharCount)

	fmt.Println("How many numbers?")
	fmt.Scanln(&input.NumberCount)

	if !checkCondition(input) {
		fmt.Println("Can not generate !")
		return
	}

	output := Practice37Normal(input)
	fmt.Println("Your password is: ", output.Password)

}

func Practice37Normal(input Input) (output Output) {

	output = implementP037(input)

	if !isLegalPassword(input, output.Password) {
		fmt.Println("Not Legal")
		return
	}

	return
}
