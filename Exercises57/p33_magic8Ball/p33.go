package p33

import (
	"bufio"
	"fmt"
	"os"
)

type Input struct {
}

type Output struct {
}

func RunP33() {

	inputReader1 := bufio.NewReader(os.Stdin)

	fmt.Println("What's your question?")
	_, errOsInput := inputReader1.ReadString('\n')

	if errOsInput != nil {
		fmt.Println(" Something wrong", errOsInput)
	}

	s := getLuckyBall()
	fmt.Println(s)

}

func Practice33Normal(input Input) (output Output) {

	return
}
