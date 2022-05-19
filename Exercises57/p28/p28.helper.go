package p28

func getOrderNumberUnit(n int) (s string) {

	switch n {
	case 1:
		s = "-st"
	case 2:
		s = "-nd"
	case 3:
		s = "-rd"
	default:
		s = "-th"
	}

	return

}
