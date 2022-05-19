package p20

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Input struct {
	Amount float64
	State  string
	Town   string
}

type Output struct {
	Tax   float64
	Total float64
}

func RunP20() {

	var input Input
	var errOsInput error
	var s1 string
	inputReader1 := bufio.NewReader(os.Stdin)

	fmt.Println("What is the order amount?")
	fmt.Scanln(&input.Amount)

	fmt.Println("What state do you live in?")
	fmt.Scanln(&input.State)
	//input.State = strings.ToLower(input.State)

	if isInWisconsinState(input.State) {
		fmt.Println("What town do you live in?")
		s1, errOsInput = inputReader1.ReadString('\n')

		if errOsInput != nil {
			fmt.Println(" wrong town input data", errOsInput)
			return
		}

		s1 = strings.TrimRight(s1, "\n")
		input.Town = trimLastChar(s1) //最後一個字元會有空白

		//fmt.Scanln(&input.Town) //<---- Need to revise
		//input.Town = strings.ToLower(input.Town)
	}

	output := Practice20Normal(input)

	if output.Tax > 0 {
		fmt.Println("The tax is $", output.Tax, ".")
	}

	fmt.Println("The total is $", output.Total, ".")

}

func Practice20Normal(input Input) (output Output) {

	output.Tax = calTax(input)
	output.Total = input.Amount + output.Tax

	return
}
