package p17

import (
	"fmt"
)

type Input struct {
	Amount float64
	Weight float64
	Hour   float64
	Sex    int
}

type Output struct {
	BAC    float64
	Result string
}

func RunP17() {

	var input Input

	fmt.Println("How much did you drink? (unit: ounce)")
	fmt.Scanln(&input.Amount)

	fmt.Println("How much do you weigh? (unit: lb)")
	fmt.Scanln(&input.Weight)

	fmt.Println("How long since you drink? (unit: hour)")
	fmt.Scanln(&input.Hour)

	fmt.Println("What is your sex ? (man: 1,woman:0)")
	fmt.Scanln(&input.Sex)

	if !isLegalSex(input.Sex) {
		fmt.Println(" Wrong sex input")
		return
	}

	output := Practice17Normal(input)
	fmt.Println("Your BAC is ", output.BAC)
	fmt.Println(output.Result)

}

func Practice17Normal(input Input) (output Output) {

	r := getRateSex(input.Sex)

	output.BAC = evaluateBAC(input, r)

	if isLegalBAC(output.BAC) {
		output.Result = ResultLegal
	} else {
		output.Result = ResultNotLegal
	}

	return
}
