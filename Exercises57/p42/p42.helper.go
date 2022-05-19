package p42

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetData() (input Input) {

	filePath, _ := filepath.Abs("../Practices57/p42/" + "p42Raw.txt")

	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			input.RawData = append(input.RawData, scanner.Text())
		}
	}

	return
}

func ProcessData(input Input) (output Output) {

	for _, rawData := range input.RawData {

		//fmt.Println("rawData= ", rawData)

		rs := strings.Replace(rawData, ",", " ", -1)

		//fmt.Println("rs= ", rs)

		output.ProcessedData = append(output.ProcessedData, rs)

	}

	return

}

func OutputTheProcessedData(output Output) {

	filePath, _ := filepath.Abs("../Practices57/p42/" + "p42Processed.txt")

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("os Create error: ", err)
		return
	}
	defer f.Close()

	headString := "Last First Salary"
	separatorBarString := "------------"

	bw := bufio.NewWriter(f)

	bw.WriteString(headString + "\n")
	bw.WriteString(separatorBarString + "\n")

	for _, data := range output.ProcessedData {
		bw.WriteString(data + "\n")
	}
	bw.Flush()

}
