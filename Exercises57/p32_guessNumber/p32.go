package p32

import (
	"fmt"
	"strconv"
)

type Input struct {
}

type Output struct {
}

func RunP32() {

	var s string
	var level, goal, value int
	var err error
	var count int

	fmt.Println("Let's play Guess the Number.")

	for {
		fmt.Println("Pick a difficulty level (1,2, or 3):")
		fmt.Scanln(&s)

		level, err = strconv.Atoi(s)

		if err != nil {
			fmt.Println("convert input to int fail!,error ", err)
			continue
		}

		if !isLegalLevel(level) {
			fmt.Println("No such level", level)
			continue
		}

		break
	}

	goal = getGaolNumber(level)

	fmt.Println("I have my number. What's your guess?")
	fmt.Scanln(&s)
	count++

	for {

		value, err = strconv.Atoi(s)

		if err != nil {
			fmt.Println("convert input to int fail!,error ", err)
			continue
		}

		if value == goal {
			countUnit := getCountUnitSuffix(count)
			fmt.Println("You got it in", count, countUnit)
			break
		}

		if value < goal {
			fmt.Println("Too low!")
		}

		if value > goal {
			fmt.Println("Too high!")
		}

		fmt.Println("Guess again:")
		fmt.Scanln(&s)
		count++

	}

	fmt.Println("Play again?  (y/n)")
	fmt.Scanln(&s)

	if willPlayAgain(s) {
		RunP32()
	}

}

func Practice32Normal(input Input) (output Output) {

	return
}
