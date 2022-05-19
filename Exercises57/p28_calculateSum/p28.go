package p28

import (
	"fmt"
)

type Input struct {
	Counter int
	Numbers []int
}

type Output struct {
	Total int
}

func RunP28() {

	var input Input
	var n int

	fmt.Println("How many numbers you want to sum up? ")
	fmt.Scanln(&input.Counter)

	for i := 0; i < input.Counter; i++ {

		val := i + 1
		unit := getOrderNumberUnit(val)
		fmt.Println("Please enter the ", val, unit, " number:")
		fmt.Scanln(&n)
		input.Numbers = append(input.Numbers, n)

	}

	output := Practice28Normal(input)

	fmt.Println("The total is ", output.Total)

}

func Practice28Normal(input Input) (output Output) {

	for _, v := range input.Numbers {
		output.Total += v
	}

	return
}
