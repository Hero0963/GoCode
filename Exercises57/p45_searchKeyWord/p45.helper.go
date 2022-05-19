package p45

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateNewData(oriStr string, repStr string) {

	filePathNew, _ := filepath.Abs("../Practices57/p45/" + "p45New.txt")

	f, err := os.Create(filePathNew)
	if err != nil {
		fmt.Println("os Create error: ", err)
		return
	}
	defer f.Close()

	bw := bufio.NewWriter(f)

	filePathRaw, _ := filepath.Abs("../Practices57/p45/" + "p45Raw.txt")

	if file, err := os.Open(filePathRaw); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {

			s := scanner.Text()
			rs := strings.Replace(s, oriStr, repStr, -1)

			bw.WriteString(rs + "\n")
			bw.Flush()

		}
	}

	return
}
