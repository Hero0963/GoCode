package p22

func getOrdinalNumbersSuffix(num int) (s string) {

	switch num {
	case 1:
		s = "st"
	case 2:
		s = "nd"
	case 3:
		s = "rd"
	default:
		s = "th"

	}

	return s

}
