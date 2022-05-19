package p03

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func RunP03() {

	var s1, s2 string
	var errOsInput error
	inputReader1 := bufio.NewReader(os.Stdin)
	inputReader2 := bufio.NewReader(os.Stdin)

	fmt.Println("What is the quote? ")
	s1, errOsInput = inputReader1.ReadString('\n')

	if errOsInput != nil {
		fmt.Println(" Something wrong", errOsInput)
	}

	fmt.Println("Who said it? ")
	s2, errOsInput = inputReader2.ReadString('\n')
	if errOsInput != nil {
		fmt.Println(" Something wrong", errOsInput)
	}

	s1 = strings.TrimRight(s1, "\n")
	s2 = strings.TrimRight(s2, "\n")
	s1 = trimLastChar(s1) //最後一個字元會有空白
	s2 = trimLastChar(s2)

	// s1 = trimLastChar(s1)
	// s2 = trimLastChar(s2)

	// fmt.Println("s1= ", s1)
	// fmt.Println("s2= ", s2)

	r := Practice03Normal(s1, s2)

	fmt.Println(r)

}

func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}

func combineTheInput(si []string) (so string) {

	for i, v := range si {
		if i == 0 {
			so += v
		} else {
			so += (" " + v)
		}
	}

	return so
}

func Practice03Normal(s1 string, s2 string) (r string) {

	// fmt.Println(len(s2))
	// fmt.Println(len(s1))

	// bs1 := []byte(s1)
	// fmt.Println(bs1)

	// for _, v := range bs1 {
	// 	fmt.Println(string(v))
	// }

	// bs2 := []byte(s2)
	// fmt.Println(bs2)

	// for _, v := range bs2 {
	// 	fmt.Println(string(v))
	// }

	r = s2 + StrPrefix + s1

	return r

}
