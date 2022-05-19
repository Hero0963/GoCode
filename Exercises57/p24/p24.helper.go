package p24

func AreAnagrams(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	bs1 := []byte(s1)
	bs2 := []byte(s2)

	m := make(map[byte]int)

	for _, v := range bs1 {
		m[v]++
	}

	for _, val := range bs2 {
		m[val]--
	}

	for _, value := range m {
		if value != 0 {
			return false
		}

	}

	return true

}
