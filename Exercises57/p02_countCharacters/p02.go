package p02

import (
	"fmt"
	"strconv"
)

func RunP02() {

	var s string
	fmt.Println("What is the input string? ")
	fmt.Scanln(&s)

	r := Practice02Normal(s)

	fmt.Println(r)

}

func Practice02Normal(s string) (r string) {

	l := len(s)
	ss := strconv.Itoa(l)

	r = s + StrPrefix + ss + StrSuffix

	return r

}
