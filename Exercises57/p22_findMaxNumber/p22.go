package p22

import (
	"fmt"
	"sort"
)

type Input struct {
	Numbers []int
}

type Output struct {
	Numbers []int
}

func RunP22() {

	var input Input
	var goal int
	var tmp int
	var onSuffix string

	fmt.Println("How many numbers do you want to compare?")
	fmt.Scanln(&goal)

	for i := 0; i < goal; i++ {
		onSuffix = getOrdinalNumbersSuffix(i + 1)
		fmt.Println("Please enter the ", i+1, onSuffix, " number: ")
		fmt.Scanln(&tmp)
		input.Numbers = append(input.Numbers, tmp)

	}

	output := Practice22Normal(input)

	fmt.Println("The result is: ", output.Numbers)

}

func Practice22Normal(input Input) (output Output) {

	sort.Ints(input.Numbers)
	output.Numbers = input.Numbers

	return
}
