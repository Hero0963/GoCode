package p41

import (
	"fmt"
	"sort"
)

type Input struct {
	RawData []string
}

type Output struct {
	SortedData []string
}

func RunP41() {

	fmt.Println("This program will read the txt file and output the sorted data : ")

	input := GetData()
	output := Practice41Normal(input)

	OutputTheSortedData(output)

}

func Practice41Normal(input Input) (output Output) {

	ss := input.RawData
	sort.Strings(ss)
	output.SortedData = ss

	return
}
