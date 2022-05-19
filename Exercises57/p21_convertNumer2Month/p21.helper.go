package p21

func isLegalNumberInput(num int) bool {

	if num < NumberLowerBound {
		return false
	}

	if num > NumberUpperBound {
		return false
	}

	return true
}

func getNumberToNameMap() map[int]string {
	m := make(map[int]string)

	m[1] = "January"
	m[2] = "February"
	m[3] = "March"
	m[4] = "April"
	m[5] = "May"
	m[6] = "June"
	m[7] = "July"
	m[8] = "August"
	m[9] = "September"
	m[10] = "October"
	m[11] = "November"
	m[12] = "December"

	return m

}
