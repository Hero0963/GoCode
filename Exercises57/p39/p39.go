package p39

import (
	"fmt"
)

type Input struct {
	FirstName      string
	LastName       string
	Position       string
	SeparationDate string
}

type Output struct {
	FullName       string
	Position       string
	SeparationDate string
}

func RunP39() {

	fmt.Println("This program will show the employee list sort by last name: ")

	input := GetData()

	output := Practice39Normal(input)

	fmt.Println("Name----Position----Separation: ")

	for _, v := range output {
		fmt.Println(v)
	}

}

func Practice39Normal(input []Input) (output []Output) {

	output = sortEmployeeListByLastName(input)

	return
}
