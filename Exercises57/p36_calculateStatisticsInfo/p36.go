package p36

import (
	"fmt"
	"strconv"
)

type Input struct {
	Numbers []float64
}

type Output struct {
	Mean float64
	Max  float64
	Min  float64
	StdD float64
}

func RunP36() {

	var input Input
	var s string

	fmt.Println("This program will calculate the mean and standard deviation.")
	fmt.Println("You can enter done to finish entering the number.")

	for {

		fmt.Println("Please enter a number:")
		fmt.Scanln(&s)

		if finishEnteringNumbers(s) {
			break
		}

		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			continue
		}

		input.Numbers = append(input.Numbers, f)

	}

	output := Practice36Normal(input)
	fmt.Println("numbers:", input.Numbers)
	fmt.Println("The average is: ", output.Mean)
	fmt.Println("The max is: ", output.Max)
	fmt.Println("The min is: ", output.Min)
	fmt.Println("The standard deviation is: ", output.StdD)

}

func Practice36Normal(input Input) (output Output) {

	output.Mean = getMean(input.Numbers)
	output.Max = getMaxNumber(input.Numbers)
	output.Min = getMinNumber(input.Numbers)
	output.StdD = getStandardDeviation(input.Numbers)

	return
}
