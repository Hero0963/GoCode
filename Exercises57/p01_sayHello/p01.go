package p01

import (
	"fmt"
)

func RunP01() {

	var s string
	fmt.Println("What's your name?")
	fmt.Scanln(&s)

	r := Practice01Normal(s)

	fmt.Println(r)

}

func Practice01Normal(s string) (r string) {

	r = StrPrefix + s + StrSuffix

	return r

}
