package p09

import (
	"fmt"
)

func RunP09() {

	var areaCeiling int

	fmt.Println("Enter the area of ceiling: (unit= square feet)")
	fmt.Scanln(&areaCeiling)

	numberOfGallon := Practice09Normal(areaCeiling)
	gallonUnit := getGallonUnit(numberOfGallon)

	fmt.Println("You will need to purchase ", numberOfGallon, gallonUnit, " of paint to cover ", areaCeiling, " square feet.")
}

func Practice09Normal(areaCeiling int) (numberOfGallon int) {
	numberOfGallon = areaCeiling / Converter
	r := areaCeiling % Converter

	if r > 0 {
		numberOfGallon++
	}

	return numberOfGallon

}

func getGallonUnit(numberOfGallon int) (gallonUnit string) {

	gallonUnit = " gallon"

	if numberOfGallon > 1 {
		gallonUnit += "s"
	}

	return gallonUnit
}
