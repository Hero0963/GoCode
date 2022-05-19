package p25

func isAVeryWeakPassword(s string) bool {

	if len(s) > 8 {
		return false
	}

	bs := []byte(s)
	for _, v := range bs {
		if !isANumberInByteForm(v) {
			return false
		}
	}

	return true
}

func isAWeakPassword(s string) bool {

	if len(s) > 8 {
		return false
	}

	bs := []byte(s)
	for _, v := range bs {
		if !isAnAlphaBetInByteForm(v) {
			return false
		}
	}

	return true
}

func isAStrongPassword(s string) bool {

	if len(s) < 8 {
		return false
	}

	var countAlphabet, countNumber int

	bs := []byte(s)
	for _, v := range bs {

		if isAnAlphaBetInByteForm(v) {
			countAlphabet++
		} else if isANumberInByteForm(v) {
			countNumber++
		}
	}

	if countAlphabet == 0 || countNumber == 0 {
		return false
	}

	return true
}

func isAVeryStrongPassword(s string) bool {

	if len(s) < 8 {
		return false
	}

	var countAlphabet, countNumber, countSpecialChar int

	bs := []byte(s)
	for _, v := range bs {

		if isAnAlphaBetInByteForm(v) {
			countAlphabet++
		} else if isANumberInByteForm(v) {
			countNumber++
		} else {
			countSpecialChar++
		}
	}

	if countAlphabet == 0 || countNumber == 0 || countSpecialChar == 0 {

		//fmt.Println(s, countAlphabet, countNumber, countSpecialChar)
		return false
	}

	return true
}

func isASpecialCharInByteForm(b byte) bool {

	if !isANumberInByteForm(b) && !isAnAlphaBetInByteForm(b) {
		return true
	}

	return false
}

func isANumberInByteForm(b byte) bool {

	if b >= '1' && b <= '9' {
		return true
	}

	return false
}

func isAnAlphaBetInByteForm(b byte) bool {

	if b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' {
		return true
	}

	return false
}
