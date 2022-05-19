package p44

import (
	"fmt"
)

type Input struct {
	Name     string
	Price    float64
	Quantity int
}

type Output struct {
	Name     string
	Price    float64
	Quantity int
}

func RunP44() {

	var name string
	fmt.Println("This program will find the product's information by name: ")

	fmt.Println("What is the product name? ")
	fmt.Scanln(&name)

	data := GetData()
	output := Practice44Normal(data, name)

	if output.Name == name {
		fmt.Println("Name: ", output.Name)
		fmt.Println("Price: ", output.Price)
		fmt.Println("Quantity: ", output.Quantity)
	} else {
		fmt.Println("Sorry, that product was not found in our inventory.")
	}

}

func Practice44Normal(data []Input, name string) (output Output) {

	output = findProductByName(data, name)

	return
}
