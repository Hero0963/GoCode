package p27

import (
	"fmt"
)

type Input struct {
	FirstName  string
	LastName   string
	ZipCode    string
	EmployeeID string
}

type Output struct {
	CheckFirstName  bool
	CheckLastName   bool
	CheckZipCode    bool
	CheckEmployeeID bool
}

func RunP27() {

	var input Input

	fmt.Println("Enter the first name (at least 2 chars) :")
	fmt.Scanln(&input.FirstName)

	fmt.Println("Enter the last name (at least 2 chars) : ")
	fmt.Scanln(&input.LastName)

	fmt.Println("Enter the ZIP code (number-only):")
	fmt.Scanln(&input.ZipCode)

	fmt.Println("Enter theEmployeeID ( in the form of AA-1234 ):")
	fmt.Scanln(&input.EmployeeID)

	fmt.Println()

	output := Practice27Normal(input)
	if isCorrectData(output) {
		fmt.Println("There were no errors.")
	}

}

func Practice27Normal(input Input) (output Output) {

	output.CheckFirstName = isLegalFirstName(input.FirstName)
	output.CheckLastName = isLegalLastName(input.LastName)
	output.CheckEmployeeID = isLegalEmployeeID(input.EmployeeID)
	output.CheckZipCode = isLegalZipCode(input.ZipCode)

	return
}
