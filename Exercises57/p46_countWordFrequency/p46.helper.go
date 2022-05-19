package p46

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func GetData() (statData StatData) {

	statData = make(StatData)

	filePath, _ := filepath.Abs("../Practices57/p46/" + "p46Raw.txt")

	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {

			ss := strings.Split(scanner.Text(), " ")

			for _, str := range ss {
				statData[str]++
			}

		}
	}

	return

}

func SaveStatData(statData StatData) {

	filePathNew, _ := filepath.Abs("../Practices57/p46/" + "p46New.txt")

	f, err := os.Create(filePathNew)
	if err != nil {
		fmt.Println("os Create error: ", err)
		return
	}
	defer f.Close()

	bw := bufio.NewWriter(f)

	for key, value := range statData {

		bw.WriteString(key + ": " + strconv.Itoa(value) + "\n")
		bw.Flush()

	}

	return
}
