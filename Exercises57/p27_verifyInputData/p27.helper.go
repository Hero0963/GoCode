package p27

import "fmt"

func isLegalFirstName(s string) bool {

	if len(s) < LengthLB {
		fmt.Println("FirstName input is too short.")
		return false
	}

	fmt.Println("FirstName input is legal.")
	return true
}

func isLegalLastName(s string) bool {

	if len(s) < LengthLB {
		fmt.Println("LastName input is too short.")
		return false
	}

	fmt.Println("LastName input is legal.")
	return true
}

func isLegalEmployeeID(s string) bool {

	if len(s) != LengthEmployeeID {
		fmt.Println("EmployeeID input is not legal.")
		return false
	}

	bs := []byte(s)

	for i, v := range bs {
		if i < 2 {
			if !isAlphabetInByteForm(v) {
				fmt.Println("EmployeeID input is not legal.")
				return false
			}

		} else if i == 2 {

			if v != '-' {
				fmt.Println("EmployeeID input is not legal.")
				return false
			}

		} else {
			if !isNumberInByteForm(v) {
				fmt.Println("EmployeeID input is not legal.")
				return false
			}

		}

	}

	fmt.Println("EmployeeID input is legal.")
	return true

}

func isLegalZipCode(s string) bool {

	bs := []byte(s)

	for _, v := range bs {
		if !isNumberInByteForm(v) {
			fmt.Println("ZipCode input is not legal.")
			return false
		}
	}

	fmt.Println("ZipCode input is legal.")
	return true
}

func isCorrectData(output Output) bool {
	if output.CheckFirstName && output.CheckLastName && output.CheckEmployeeID && output.CheckZipCode {
		return true
	}

	return false
}

func isAlphabetInByteForm(b byte) bool {

	if b >= 'a' && b <= 'z' {
		return true
	}

	if b >= 'A' && b <= 'Z' {
		return true
	}

	return false
}

func isNumberInByteForm(b byte) bool {

	if b >= '0' && b <= '9' {
		return true
	}

	return false
}
