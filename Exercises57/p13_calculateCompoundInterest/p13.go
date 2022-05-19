package p13

import (
	"fmt"
)

type Input struct {
	Principal     float64
	Rate          float64
	NumberOfYears int
	InterestTimes int
}

type Output struct {
	Worth float64
}

func RunP13() {

	var input Input

	fmt.Println("What is the principal amount?")
	fmt.Scanln(&input.Principal)

	fmt.Println("What is the rate?")
	fmt.Scanln(&input.Rate)

	fmt.Println("What is the number of years?")
	fmt.Scanln(&input.NumberOfYears)

	fmt.Println("What is the number of times the interest is compounded per year?")
	fmt.Scanln(&input.InterestTimes)

	output := Practice13Normal(input)

	fmt.Println(" the investment will be worth $", output.Worth, " .")

}

func Practice13Normal(input Input) (output Output) {

	totalRound := input.NumberOfYears * input.InterestTimes

	var r float64
	r = 1.0 + (input.Rate/100.0)/float64(input.InterestTimes)

	output.Worth = input.Principal

	for i := 0; i < totalRound; i++ {
		output.Worth *= r
	}

	tmp := int(output.Worth * 100)

	output.Worth = float64(tmp) / 100.0

	return output

}
