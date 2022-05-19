package p44

func GetData() (products []Input) {

	var input Input

	input.Name = "Widget"
	input.Price = 25.00
	input.Quantity = 5

	products = append(products, input)

	input.Name = "Thing"
	input.Price = 15.00
	input.Quantity = 5

	products = append(products, input)

	input.Name = "Doodad"
	input.Price = 5.00
	input.Quantity = 10

	products = append(products, input)

	return
}

func findProductByName(products []Input, name string) (output Output) {

	for _, product := range products {
		if product.Name == name {
			output = Output(product)
			return

		}

	}

	return
}
