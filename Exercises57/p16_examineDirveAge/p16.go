package p16

import (
	"fmt"
)

type Input struct {
	Age int
}

type Output struct {
	Result string
}

func RunP16() {

	var input Input

	fmt.Println("What is your age?")
	fmt.Scanln(&input.Age)

	output := Practice16Normal(input)
	fmt.Println(output.Result)

}

func Practice16Normal(input Input) (output Output) {

	if input.Age < 1 {
		output.Result = ResultWrongInput
	} else {

		if input.Age < LegalDriveAge {
			output.Result = ResultNotLegal
		} else {
			output.Result = ResultLegal
		}

	}

	return
}
