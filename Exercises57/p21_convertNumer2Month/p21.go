package p21

import (
	"fmt"
)

type Input struct {
	Number int
}

type Output struct {
	Name string
}

func RunP21() {

	var input Input

	fmt.Println("Please enter the number of the month:")
	fmt.Scanln(&input.Number)

	if !isLegalNumberInput(input.Number) {
		fmt.Println(" illegal input !")
		return
	}

	output := Practice21Normal(input)

	fmt.Println("The name of the month is ", output.Name, ".")

}

func Practice21Normal(input Input) (output Output) {

	m := getNumberToNameMap()

	output.Name = m[input.Number]

	return
}
