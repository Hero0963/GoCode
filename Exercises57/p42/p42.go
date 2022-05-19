package p42

import "fmt"

type Input struct {
	RawData []string
}

type Output struct {
	ProcessedData []string
}

func RunP42() {

	fmt.Println("This program will read the txt file (content is in csv form)  and output the processed data : ")

	input := GetData()
	output := Practice42Normal(input)

	OutputTheProcessedData(output)

}

func Practice42Normal(input Input) (output Output) {

	output = ProcessData(input)

	return
}
