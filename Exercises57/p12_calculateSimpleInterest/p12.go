package p12

import (
	"fmt"
)

type Input struct {
	Principal      float64
	RateOfInterest float64
	NumberOfYears  int
}

type Output struct {
	Worth float64
}

func RunP12() {

	var input Input

	fmt.Println("Enter the principal: ")
	fmt.Scanln(&input.Principal)

	fmt.Println("Enter the rate of interest: ")
	fmt.Scanln(&input.RateOfInterest)

	fmt.Println("Enter the number of years: ")
	fmt.Scanln(&input.NumberOfYears)

	output := Practice12Normal(input)

	fmt.Println(" the investment will be worth $", output.Worth, " .")

}

func Practice12Normal(input Input) (output Output) {

	output.Worth = input.Principal * (1.0 + (input.RateOfInterest/100.0)*float64(input.NumberOfYears))

	tmp := int(output.Worth * 100)

	output.Worth = float64(tmp) / 100.0

	return output

}
