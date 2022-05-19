package p26

import (
	"fmt"
	"math"
)

type Input struct {
	Balance    float64
	APR        float64
	MonthlyPay float64
}

type Output struct {
	NeedMonth int
}

func RunP26() {

	var input Input

	fmt.Println("What is your balance?")
	fmt.Scanln(&input.Balance)

	fmt.Println("What is the APR on the card (as a percent)?")
	fmt.Scanln(&input.APR)

	fmt.Println("What is the monthly payment you can make?")
	fmt.Scanln(&input.MonthlyPay)

	output := Practice26Normal(input)
	MonthUnit := getMonthUnit(output.NeedMonth)

	fmt.Println("It wil take you ", output.NeedMonth, " ", MonthUnit, " to pay off this card.")

}

func Practice26Normal(input Input) (output Output) {

	onePlusDayRate := float64(1) + ((input.APR / float64(100)) / float64(365))
	ph := getPowerHelper(onePlusDayRate)

	tmp := (float64(-1.0) / float64(30)) * ((math.Log(float64(1) + (input.Balance/input.MonthlyPay)*(float64(1)-ph))) /
		(math.Log(onePlusDayRate)))

	n := int(tmp)

	if tmp-float64(n) > 0 {
		n++
	}

	output.NeedMonth = n

	return
}
