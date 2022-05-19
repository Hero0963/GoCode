package p38

import "unicode/utf8"

func convertString2Numbers(s string) (numbers []int) {

	bs := []byte(s)

	for _, v := range bs {

		if isNumberByte(v) {

			n := convertByteNumber2Integer(v)
			if n != NotNum {
				numbers = append(numbers, n)
			}

		}

	}

	return numbers
}

func filterEvenNumbers(oriNumbers []int) (EvenNumbers []int) {

	for _, v := range oriNumbers {

		if v%2 == 0 {
			EvenNumbers = append(EvenNumbers, v)
		}

	}

	return
}

func convertByteNumber2Integer(b byte) int {

	if !isNumberByte(b) {
		return NotNum
	}

	return int(b - '0')

}

func isNumberByte(b byte) bool {

	if b >= '0' && b <= '9' {
		return true
	}

	return false
}

func isAEvenNumberByte(b byte) bool {

	evenList := getEvenNumByteList()

	for _, v := range evenList {
		if v == b {
			return true
		}
	}

	return false
}

func getNumByteList() []byte {
	return []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
}

func getEvenNumByteList() []byte {
	return []byte{'0', '2', '4', '6', '8'}
}

func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}
