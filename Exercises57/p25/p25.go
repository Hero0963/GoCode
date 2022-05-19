package p25

import (
	"fmt"
)

type Input struct {
	Password string
}

type Output struct {
	Result string
}

func RunP25() {

	var input Input

	fmt.Println("Enter your password, we will tell you how safe is your passwrd: ")
	fmt.Scanln(&input.Password)

	output := Practice25Normal(input)
	fmt.Println(output.Result)

}

func Practice25Normal(input Input) (output Output) {

	if isAVeryStrongPassword(input.Password) {
		output.Result = ResultVeryStrong
		return
	}

	if isAStrongPassword(input.Password) {
		output.Result = ResultStrong
		return
	}

	if isAWeakPassword(input.Password) {
		output.Result = ResultWeak
		return
	}

	if isAVeryWeakPassword(input.Password) {
		output.Result = ResultVeryWeak
		return
	}

	return
}
