package p24

import (
	"fmt"
)

type Input struct {
	StrFisrt  string
	StrSecond string
}

type Output struct {
	Result bool
}

func RunP24() {

	var input Input

	fmt.Println("Enter two strings and I'll tell you if they are anagrams.")

	fmt.Println("Enter the first string: ")
	fmt.Scanln(&input.StrFisrt)

	fmt.Println("Enter the second string: ")
	fmt.Scanln(&input.StrSecond)

	output := Practice24Normal(input)

	if output.Result {
		fmt.Println("They are anagrams.")
	} else {
		fmt.Println("They are not anagrams.")
	}

}

func Practice24Normal(input Input) (output Output) {

	output.Result = AreAnagrams(input.StrFisrt, input.StrSecond)

	return
}
