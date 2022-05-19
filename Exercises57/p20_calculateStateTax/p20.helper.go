package p20

import (
	"strings"
	"unicode/utf8"
)

func isInTaxStates(s string) bool {
	ss := getTaxStatesList()

	for _, v := range ss {
		if v == s {
			return true
		}
	}

	return false
}

func getTaxStatesList() []string {

	s := []string{"Wisconsin", "Illinois"}

	return s
}

func isInWisconsinState(s string) bool {

	if sameStringIgnoreUpperLowerCase(s, "Wisconsin") {
		return true
	}

	return false

}

func isInIllinoisState(s string) bool {

	if sameStringIgnoreUpperLowerCase(s, "Illinois") {
		return true
	}

	return false

}

func isInWExtraTown(s string) bool {

	ss := getWExtraTownList()

	for _, v := range ss {

		if sameStringIgnoreUpperLowerCase(s, v) {
			return true
		}

	}

	return false

}

func getWExtraTownList() []string {

	s := []string{"Eau Claire", "Dunn Country"}

	return s
}

func sameStringIgnoreUpperLowerCase(s1 string, s2 string) bool {

	ls1 := strings.ToLower(s1)
	ls2 := strings.ToLower(s2)

	if ls1 == ls2 {
		return true
	}

	return false

}

func calTax(input Input) (tax float64) {
	if isInTaxStates(input.State) {

		if isInIllinoisState(input.State) {
			tax = input.Amount * TaxStateIllinois
			return
		}

		if isInWisconsinState(input.State) && isInWExtraTown(input.Town) {

			tax = calWExtraTaxTown(input)
			return
		}

	}

	return

}

func calWExtraTaxTown(input Input) (tax float64) {

	ss := getWExtraTownList()

	for _, v := range ss {
		if sameStringIgnoreUpperLowerCase(input.Town, v) {

			r := getWExtraTaxTownRatio(input.Town)
			//fmt.Println("r=", r)

			tax = input.Amount * r
			return
		}

	}

	return
}

func getWExtraTaxTownRatio(s string) float64 {

	ls := strings.ToLower(s)

	m := getWExtraTaxTownRatioMap()

	return m[ls]

}

func getWExtraTaxTownRatioMap() map[string]float64 {

	m := make(map[string]float64)

	m["eau claire"] = TaxTownEauClarie
	m["dunn country"] = TaxTownDunnCountry

	return m
}

func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}
