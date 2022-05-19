package p10

import (
	"fmt"
)

type Item struct {
	Price    float64
	Quantity int
}

type CalPrice struct {
	SubTotal float64
	Tax      float64
	Total    float64
}

func RunP10() {

	var items []Item
	var count int

	for count < 3 {

		var item Item

		fmt.Println("Enter the price of item ", (count + 1), " :")
		fmt.Scanln(&item.Price)

		fmt.Println("Enter the quantity of item ", (count + 1), " :")
		fmt.Scanln(&item.Quantity)

		items = append(items, item)

		count++

	}

	calPrice := Practice10Normal(items)

	fmt.Println(" Subtotal: $", calPrice.SubTotal)
	fmt.Println(" Tax: $", calPrice.Tax)
	fmt.Println(" Total: $", calPrice.Total)

}

func Practice10Normal(items []Item) (calPrice CalPrice) {

	for _, v := range items {
		calPrice.SubTotal += v.Price * float64(v.Quantity)
	}

	calPrice.Tax = calPrice.SubTotal * TaxConverter

	calPrice.Total = calPrice.SubTotal + calPrice.Tax

	return calPrice

}
