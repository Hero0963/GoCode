package p29

import (
	"fmt"
	"strconv"
)

func calRuleOf72(input Input) (output Output) {

	n, err := strconv.Atoi(input.Rate)

	if err != nil {
		fmt.Println("Convert input data to int error!!")
		fmt.Println("err= ", err)
		fmt.Println("rate= ", input.Rate)
		return
	}

	year := int(float64(72) / float64(n))

	if float64(72)-float64(year)*float64(n) > 0 {
		year++
	}

	output.Year = year

	return

}

func isInputLeagal(input Input) bool {

	if isNotZeroValue(input) && isANumber(input) {
		return true
	}

	return false

}

func isNotZeroValue(input Input) bool {

	bs := []byte(input.Rate)

	for _, v := range bs {
		if v != '0' {
			return true
		}
	}

	fmt.Println("input can not be 0.")
	return false

}

func isANumber(input Input) bool {

	bs := []byte(input.Rate)

	for _, v := range bs {

		if v >= '0' && v <= '9' {
			continue
		}

		fmt.Println("It is not a valid number.")
		return false

	}

	return true
}

func getYearUnit(n int) (s string) {

	s = " year"

	if n > 1 {
		s += "s"
	}

	return s

}
