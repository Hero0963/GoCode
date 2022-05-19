package p34

import (
	"fmt"
	"unicode/utf8"
)

type EmployeeList []string

func getDefaultList() EmployeeList {
	eL := EmployeeList{"John Smith", "Jackie Jackson", "Chris Jones", "Amanda Cullen", "Jeremy Goodwin"}
	return eL
}

func (eList EmployeeList) ShowList() {

	fmt.Println("The employee list is the following:")

	for _, v := range eList {
		fmt.Println(v)
	}
}

func (eList EmployeeList) Delete(name string) EmployeeList {

	rL := EmployeeList{}

	for _, v := range eList {
		if v != name {
			rL = append(rL, v)
		}
	}

	return rL
}

func (eList EmployeeList) In(name string) bool {

	for _, v := range eList {
		if v == name {
			return true
		}
	}

	return false
}

func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}
