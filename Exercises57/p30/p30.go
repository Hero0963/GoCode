package p30

import "fmt"

type Input struct {
}

type Output struct {
}

func RunP30() {

	for i := 0; i < 13; i++ {
		for j := 0; j < 13; j++ {
			fmt.Println(i, "*", j, "=", i*j)
		}
	}
}

func Practice30Normal(input Input) (output Output) {

	return
}
