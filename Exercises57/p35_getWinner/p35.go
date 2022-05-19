package p35

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Input struct {
	Names []string
}

type Output struct {
	Winner string
}

func RunP35() {

	var input Input
	var name string
	var errOsInput error
	inputReader1 := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter the competitors, this program will return the winner.")
	fmt.Println("You can enter a blank line to finish entering the name.")

	for {

		fmt.Println("Please enter a name:")
		name, errOsInput = inputReader1.ReadString('\n')

		if errOsInput != nil {
			fmt.Println(" Something wrong", errOsInput)
		}

		name = strings.TrimRight(name, "\n")
		name = trimLastChar(name) //最後一個字元會有空白

		if isAValidInput(name) {
			input.Names = append(input.Names, name)
			continue
		}

		fmt.Println("Finish entering the name.")
		break
	}

	output := Practice35Normal(input)
	fmt.Println("The winner is ...", output.Winner, ".")

}

func Practice35Normal(input Input) (output Output) {

	l := len(input.Names)

	if l == 0 {
		return
	}

	//fmt.Println(input.Names)

	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(l)

	output.Winner = input.Names[idx]

	return
}
