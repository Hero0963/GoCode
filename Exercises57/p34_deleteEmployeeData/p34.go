package p34

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Input struct {
}

type Output struct {
}

func RunP34() {

	var name string
	var errOsInput error
	inputReader1 := bufio.NewReader(os.Stdin)

	eL := getDefaultList()
	eL.ShowList()

	fmt.Println("Enter a rmployee name to remove:")
	name, errOsInput = inputReader1.ReadString('\n')

	if errOsInput != nil {
		fmt.Println(" Something wrong", errOsInput)
	}

	name = strings.TrimRight(name, "\n")
	name = trimLastChar(name) //最後一個字元會有空白

	if !eL.In(name) {
		fmt.Println("No such one.")
		return
	}

	eL = eL.Delete(name)
	eL.ShowList()

}

func Practice34Normal(input Input) (output Output) {

	return
}
