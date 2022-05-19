package p38

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Input struct {
	RawData string
}

type Output struct {
	EvenNumbers []int
}

func RunP38() {

	var input Input
	var s string
	var errOsInput error
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a list of numbers,separated by spaces: ")
	s, errOsInput = inputReader.ReadString('\n')

	if errOsInput != nil {
		fmt.Println(" Something wrong", errOsInput)
	}

	s = strings.TrimRight(s, "\n")

	s = trimLastChar(s) //最後一個字元會有空白

	input.RawData = s

	output := Practice38Normal(input)

	fmt.Println("The even numbers are: ", output.EvenNumbers)

}

func Practice38Normal(input Input) (output Output) {

	numbers := convertString2Numbers(input.RawData)
	output.EvenNumbers = filterEvenNumbers(numbers)

	return
}
