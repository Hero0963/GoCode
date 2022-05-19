package p15

import (
	"fmt"
)

type Input struct {
	UserID   string
	Password string
}

type Output struct {
	Result string
}

func RunP15() {

	var input Input

	fmt.Println("What is the user_id?")
	fmt.Scanln(&input.UserID)

	fmt.Println("What is the password?")
	fmt.Scanln(&input.Password)

	output := Practice15Normal(input)
	fmt.Println(output.Result)

}

func Practice15Normal(input Input) (output Output) {

	val, ok := UserLogin[input.UserID]

	if ok {

		if val == input.Password {
			output.Result = ResultMatch
		} else {
			output.Result = ResultWrongPWD
		}
	} else {
		output.Result = ResultNoUserID
	}

	return

}
