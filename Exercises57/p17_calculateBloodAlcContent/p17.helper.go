package p17

func evaluateBAC(input Input, r float64) (b float64) {
	b = (input.Amount*CalConst)/(input.Weight*r) - CalConst2*input.Hour

	tmp := int(b * 100)

	b = float64(tmp) / 100.0

	return
}

func isLegalSex(sex int) bool {

	if sex == SexMan {
		return true
	}

	if sex == SexWoman {
		return true
	}

	return false
}

func getRateSex(sex int) (r float64) {
	if sex == SexMan {
		r = RMan
	}
	if sex == SexWoman {
		r = RWoman
	}

	return

}

func isLegalBAC(b float64) bool {

	if b < BACLegal {
		return true
	}

	return false
}
