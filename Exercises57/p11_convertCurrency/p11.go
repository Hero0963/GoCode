package p11

import (
	"fmt"
)

type Input struct {
	Euros        float64
	ExchangeRate float64
}

type Output struct {
	USDollars float64
}

func RunP11() {

	var input Input

	fmt.Println("How manu euros are you exchanging? ")
	fmt.Scanln(&input.Euros)
	fmt.Println("What is the exchange rate? ")
	fmt.Scanln(&input.ExchangeRate)

	output := Practice11Normal(input)

	fmt.Println(" you can exchage ", output.USDollars, " U.S. dollars")

}

func Practice11Normal(input Input) (output Output) {

	output.USDollars = (input.Euros * input.ExchangeRate) / 100.0

	tmp := int(output.USDollars * 100)

	output.USDollars = float64(tmp) / 100.0

	return output

}
