package p07

import (
	"fmt"
)

func RunP07() {

	var length, width float64

	fmt.Println("What is the length of the room in feet?")
	fmt.Scanln(&length)

	fmt.Println("What is the width of the room in feet?")
	fmt.Scanln(&width)

	fmt.Println("You entered dimensions of ", length, " feet by ", width, " feet.")
	fmt.Println("The area is")

	var areaSquareFeet, areaSquareMeters float64
	areaSquareFeet = length * width
	areaSquareMeters = Practice07Normal(areaSquareFeet)

	fmt.Println(areaSquareFeet, " square feet ")
	fmt.Println(areaSquareMeters, " square meters ")

}

func Practice07Normal(areaSquareFeet float64) (areaSquareMeters float64) {
	areaSquareMeters = areaSquareFeet * AreaConverterF2M

	return areaSquareMeters

}
