package p45

import (
	"fmt"
)

type Input struct {
}

type Output struct {
}

func RunP45() {

	var oriStr string
	var repStr string
	fmt.Println("This program will replace the string to another: ")

	fmt.Println("What is the original string? ")
	fmt.Scanln(&oriStr)

	fmt.Println("What is the replaced string? ")
	fmt.Scanln(&repStr)

	Practice45Normal(oriStr, repStr)

	fmt.Println("Finished.")

}

func Practice45Normal(oriStr string, repStr string) {

	GenerateNewData(oriStr, repStr)

	return
}
