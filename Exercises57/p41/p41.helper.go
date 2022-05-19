package p41

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func GetData() (input Input) {

	filePath, _ := filepath.Abs("../Practices57/p41/" + "p41Raw.txt")

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

func OutputTheSortedData(output Output) {

	filePath, _ := filepath.Abs("../Practices57/p41/" + "p41Sorted.txt")

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("os Create error: ", err)
		return
	}
	defer f.Close()

	l := len(output.SortedData)

	headString := "Total of " + strconv.Itoa(l) + " names"
	separatorBarString := "------------"

	bw := bufio.NewWriter(f)

	bw.WriteString(headString + "\n")
	bw.WriteString(separatorBarString + "\n")

	for _, v := range output.SortedData {
		bw.WriteString(v + "\n")
	}
	bw.Flush()

}
