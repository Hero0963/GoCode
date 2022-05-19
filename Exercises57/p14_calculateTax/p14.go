package p14

import (
	"fmt"
	"strings"
)

type Input struct {
	Amount float64
	State  string
}

type Output struct {
	SubTotal float64
	Tax      float64
	Total    float64
}

func RunP14() {

	var input Input

	fmt.Println("What is the order amount?")
	fmt.Scanln(&input.Amount)

	fmt.Println("What is the state?")
	fmt.Scanln(&input.State)

	output := Practice14Normal(input)

	if IsIntheSpecialStates(input.State) {
		fmt.Println("The subtotal is $", output.SubTotal)
		fmt.Println("The tax is $", output.Tax)
	}

	fmt.Println("The total is $", output.Total)

}

func Practice14Normal(input Input) (output Output) {

	output.SubTotal = input.Amount

	if IsIntheSpecialStates(input.State) {
		output.Tax = output.SubTotal * (ExtraTax / 100.0)
	}

	output.Total = (output.SubTotal + output.Tax)

	tmp := int(output.Total * 100)

	output.Total = float64(tmp) / 100.0

	return output

}

func IsIntheSpecialStates(s string) bool {
	ss := getSpecialStates()
	sUppercase := strings.ToUpper(s)

	for _, v := range ss {
		if sUppercase == v {
			return true
		}
	}

	return false
}

func getSpecialStates() []string {
	return []string{"WI"}
}
