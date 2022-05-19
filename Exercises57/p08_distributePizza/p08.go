package p08

import (
	"fmt"
)

func RunP08() {

	var numberOfPeople, numberOfPizza int

	fmt.Println("How many people?")
	fmt.Scanln(&numberOfPeople)

	fmt.Println("How many pizzas do you have?")
	fmt.Scanln(&numberOfPizza)

	strPizzaUnit := getPizzaUnit(numberOfPizza)
	fmt.Println(numberOfPeople, " people with ", numberOfPizza, strPizzaUnit)

	numberOfShared, numberOfLeftover := Practice08Normal(numberOfPeople, numberOfPizza)

	fmt.Println("Each people gets ", numberOfShared, " pieces of pizza.")
	fmt.Println("There are ", numberOfLeftover, " leftover pieces.")

}

func Practice08Normal(numberOfPeople int, numberOfPizza int) (numberOfShared int, numberOfLeftover int) {
	numberOfShared = (numberOfPizza * 8) / numberOfPeople
	numberOfLeftover = (numberOfPizza * 8) - numberOfShared*numberOfPeople

	return numberOfShared, numberOfLeftover

}

func getPizzaUnit(numberOfPizza int) (strPizzaUnit string) {

	strPizzaUnit = " pizza"

	if numberOfPizza > 1 {
		strPizzaUnit += "s"
	}

	return strPizzaUnit
}
