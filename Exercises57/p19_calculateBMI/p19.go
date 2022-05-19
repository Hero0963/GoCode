package p19

import (
	"fmt"
)

type Input struct {
	Height float64
	Weight float64
}

type Output struct {
	BMI    float64
	Result string
}

func RunP19() {

	var input Input

	fmt.Println("Please enter your height: (unit: inch)")
	fmt.Scanln(&input.Height)

	fmt.Println("Please enter your weight: (unit: lb)")
	fmt.Scanln(&input.Weight)

	output := Practice19Normal(input)
	fmt.Println("Your BMI is ", output.BMI, " .")
	fmt.Println(output.Result)

}

func Practice19Normal(input Input) (output Output) {

	output.BMI = calBMI(input.Height, input.Weight)

	output.Result = checkBMIRange(output.BMI)

	return
}
