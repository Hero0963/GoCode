package p18

import (
	"fmt"
	"strings"
)

type Input struct {
	Unit        string
	Temperature float64
}

type Output struct {
	Result float64
}

func RunP18() {

	var input Input
	var s string

	fmt.Println("Press C to convert from Fahrenheit to Celsius.")
	fmt.Println("Press F to convert from Celsius to Fahrenheit.")
	fmt.Println("Your choice:")
	fmt.Scanln(&s)
	input.Unit = strings.ToUpper(s)

	if !isLegalTemperatureUnit(input.Unit) {
		fmt.Println(" Wrong input!")
		return
	}

	tu := getTemperatureUnit(input.Unit)
	dtu := getDualTemperatureUnit(input.Unit)

	fmt.Println("Please enter the temperature in ", dtu, " :")
	fmt.Scanln(&input.Temperature)

	output := Practice18Normal(input)
	fmt.Println("The temperature in ", tu, " is ", output.Result)

}

func Practice18Normal(input Input) (output Output) {

	if input.Unit == "C" {
		output.Result = convertFToC(input.Temperature)
	}

	if input.Unit == "F" {
		output.Result = convertCToF(input.Temperature)
	}

	return
}
