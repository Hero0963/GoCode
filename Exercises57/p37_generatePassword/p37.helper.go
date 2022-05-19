package p37

import (
	"fmt"
	"math/rand"
	"time"
)

func implementP037(input Input) (output Output) {

	rand.Seed(time.Now().Unix())

	start, end, count := 0, input.Len, input.SpecialCharCount+input.NumberCount
	availableSets := generateIntervalNonRepeatNumbers(start, end, count)
	//fmt.Println("availableSets=", availableSets)
	availableSets, specialCharIdx := pickIdx(availableSets, input.SpecialCharCount)
	availableSets, NumberIdx := pickIdx(availableSets, input.NumberCount)

	//fmt.Println("specialCharIdx=", specialCharIdx)
	//fmt.Println("NumberIdx=", NumberIdx)

	if len(availableSets) != 0 {
		fmt.Println("something wrong! availableSets is not empty! ", availableSets)
	}

	rMap := getRuleMap(specialCharIdx, NumberIdx)
	//fmt.Println("rMap=", rMap)

	output.Password = generateRandomStringByRuleMap(input.Len, rMap)

	return
}

func generateRandomStringByRuleMap(length int, rMap map[int]string) string {

	bs := make([]byte, length)

	for i := 0; i < length; i++ {
		bs[i] = getByteByRule(rMap[i])
	}

	return string(bs)
}

func getByteByRule(s string) (b byte) {

	switch s {
	case strSpecialChar:
		b = getRandomSpecialCharByte()
	case strNumber:
		b = getRandomNumberByte()
	default:
		b = getRandomAlphabetByte()
	}

	return b

}

func getRandomAlphabetByte() byte {

	start := int('a')
	end := int('z' + 1)

	//rSeed := rand.New(rand.NewSource(time.Now().UnixNano()))

	val := rand.Intn((end - start)) + start

	return byte(val)

}

func getRandomNumberByte() byte {

	start := int('0')
	end := int('9' + 1)

	//rSeed := rand.New(rand.NewSource(time.Now().UnixNano()))

	val := rand.Intn((end - start)) + start

	return byte(val)

}

func getRandomSpecialCharByte() byte {

	start := int('!')
	end := int('/' + 1)

	//rSeed := rand.New(rand.NewSource(time.Now().UnixNano()))

	val := rand.Intn((end - start)) + start

	return byte(val)

}

func getRuleMap(specialCharIdx []int, NumberIdx []int) map[int]string {

	m := make(map[int]string)

	for _, v := range specialCharIdx {
		m[v] = strSpecialChar
	}

	for _, val := range NumberIdx {
		m[val] = strNumber
	}

	return m
}

func checkCondition(input Input) bool {

	if input.Len <= 0 || input.SpecialCharCount < 0 || input.NumberCount < 0 {
		return false
	}

	if input.Len < (input.SpecialCharCount + input.NumberCount) {
		return false
	}

	return true
}

func pickIdx(before []int, number int) (after []int, picked []int) {
	if !canPickIdx(before, number) {
		return before, picked
	}

	start, end := 0, len(before)
	rNumbers := generateIntervalNonRepeatNumbers(start, end, number)
	m := getPcikedNumberMap(rNumbers)

	for index, value := range before {

		if m[index] {
			picked = append(picked, value)
		} else {
			after = append(after, value)
		}
	}

	return
}

func getPcikedNumberMap(picked []int) map[int]bool {
	m := make(map[int]bool)

	for _, v := range picked {
		m[v] = true
	}

	return m

}

func canPickIdx(before []int, number int) bool {

	if number <= 0 || number > len(before) {
		return false
	}

	return true

}

func generateIntervalNonRepeatNumbers(start int, end int, count int) (result []int) {

	//fmt.Println("generateIntervalNonRepeatNumbers input  start,end,count=", start, end, count)

	if !canGenerateRandomNumber(start, end, count) {
		fmt.Println("can not generate interval non-repeat numbers")
		return
	}

	//rSeed := rand.New(rand.NewSource(time.Now().UnixNano()))

	for len(result) < count {

		num := rand.Intn((end - start)) + start

		exist := false
		for _, v := range result {

			if v == num {
				exist = true
				break
			}

		}

		if !exist {
			result = append(result, num)
		}

	}

	return result
}

func canGenerateRandomNumber(start int, end int, count int) bool {
	if start < 0 || end <= 0 || end <= start || count > end-start+1 {
		return false
	}

	return true
}

func isLegalPassword(input Input, s string) bool {

	var cntAlphaBet, cntSpecialChar, cntNumber int
	var sType string

	bs := []byte(s)

	for _, v := range bs {

		sType = deterMineByteType(v)

		switch sType {
		case strSpecialChar:
			cntSpecialChar++
		case strNumber:
			cntNumber++
		default:
			cntAlphaBet++
		}

	}

	if input.Len == cntAlphaBet+cntSpecialChar+cntNumber && input.SpecialCharCount == cntSpecialChar && input.NumberCount == cntNumber {
		return true
	}

	return false
}

func deterMineByteType(b byte) (s string) {

	if b <= '0' && b >= '9' {
		s = strSpecialChar
		return
	}

	if b <= '!' && b >= '/' {
		s = strNumber
		return
	}

	return
}
