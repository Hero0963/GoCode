package p31

import (
	"fmt"
	"strconv"
)

type Input struct {
	Age       string
	RestingHR string
}

type Output struct {
}

func RunP31() {

	var input Input
	var age, rHR int
	var err error

	fmt.Println("This program will calculate TargetHeartRate for you:")

	for {

		fmt.Println("Please enter your age: (unit: int)")
		fmt.Scanln(&input.Age)

		age, err = strconv.Atoi(input.Age)

		if err != nil {
			fmt.Println("Convert input data to int error!!")
			fmt.Println("err= ", err)
			fmt.Println("age= ", input.Age)
			continue
		}

		break
	}

	for {
		fmt.Println("Please enter your restingHR:")
		fmt.Scanln(&input.RestingHR)

		rHR, err = strconv.Atoi(input.RestingHR)

		if err != nil {
			fmt.Println("Convert input data to int error!!")
			fmt.Println("err= ", err)
			fmt.Println("rHR= ", input.RestingHR)
			continue
		}

		break

	}

	showIntensityRateGraph(age, rHR)

}

func Practice31Normal(input Input) (output Output) {

	return
}
