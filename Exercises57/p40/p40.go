package p40

import (
	"fmt"
)

type Input struct {
	FirstName      string
	LastName       string
	Position       string
	SeparationDate string
}

type Output struct {
	FullName       string
	Position       string
	SeparationDate string
}

type ByName []Output

func (a ByName) Len() int           { return len(a) }
func (a ByName) Less(i, j int) bool { return a[i].FullName < a[j].FullName }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func RunP40() {

	var s string

	fmt.Println("Enter a search string, this program will show the employee list with name containing your search string: ")
	fmt.Scanln(&s)

	input := GetData()

	output := Practice40Normal(input, s)

	fmt.Println("Name----Position----Separation: ")

	for _, v := range output {
		fmt.Println(v)
	}

}

func Practice40Normal(input []Input, s string) (output []Output) {

	output = filterEmployeeListByName(input, s)

	return
}
