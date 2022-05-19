package p05

import "fmt"

func RunP05() {

	var firstNumber, secondNumber float64

	fmt.Println("What is the first number?")
	fmt.Scanln(&firstNumber)

	fmt.Println("What is the second number?")
	fmt.Scanln(&secondNumber)

	resultSum, resultDiff, resultProduct, resultQuotient := Practice05Normal(firstNumber, secondNumber)

	fmt.Println(firstNumber, "+", secondNumber, "=", resultSum)
	fmt.Println(firstNumber, "-", secondNumber, "=", resultDiff)
	fmt.Println(firstNumber, "*", secondNumber, "=", resultProduct)
	fmt.Println(firstNumber, "/", secondNumber, "=", resultQuotient)

}

func Practice05Normal(firstNumber float64, secondNumber float64) (sum float64, diff float64, product float64, quotient float64) {

	sum = Practice05GetSum(firstNumber, secondNumber)
	diff = Practice05GetDiff(firstNumber, secondNumber)
	product = Practice05GetProduct(firstNumber, secondNumber)
	quotient = Practice05GetQuotient(firstNumber, secondNumber)

	return
}

func Practice05GetSum(firstNumber float64, secondNumber float64) (sum float64) {

	sum = firstNumber + secondNumber
	return
}

func Practice05GetDiff(firstNumber float64, secondNumber float64) (diff float64) {

	diff = firstNumber - secondNumber
	return
}

func Practice05GetProduct(firstNumber float64, secondNumber float64) (product float64) {

	product = firstNumber * secondNumber
	return
}

func Practice05GetQuotient(firstNumber float64, secondNumber float64) (quotient float64) {

	quotient = firstNumber / secondNumber
	return
}
