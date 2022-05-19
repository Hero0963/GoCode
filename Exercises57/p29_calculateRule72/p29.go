package p29

import (
	"fmt"
)

type Input struct {
	Rate string
}

type Output struct {
	Year int
}

func RunP29() {

	var input Input

	for {

		fmt.Println("What is the rate of return?  unit:(%)")
		fmt.Scanln(&input.Rate)

		if !isInputLeagal(input) {
			fmt.Println("It is not a legal inpt data")
		} else {
			break
		}

	}

	output := Practice29Normal(input)
	unitYear := getYearUnit(output.Year)

	fmt.Println("It will take ", output.Year, unitYear, " to double your initial investment.")

}

func Practice29Normal(input Input) (output Output) {

	output = calRuleOf72(input)

	return
}
