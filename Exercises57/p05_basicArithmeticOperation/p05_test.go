package p05_test

import (
	"GoCode/Practices57/p05"
	"reflect"
	"testing"
)

type testData struct {
	first  float64
	second float64

	ans testAnswer
}

type testAnswer struct {
	sum      float64
	diff     float64
	product  float64
	quotient float64
}

func generateTestData() []testData {
	data := []testData{
		testData{first: 10, second: 5, ans: testAnswer{sum: 15, diff: 5, product: 50, quotient: 2}},
	}
	return data
}

func TestPractice05Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actualSum, actualDiff, actualProduct, actualQuotient := p05.Practice05Normal(v.first, v.second)

		var actual testAnswer
		actual.sum = actualSum
		actual.diff = actualDiff
		actual.product = actualProduct
		actual.quotient = actualQuotient

		expected := v.ans

		if reflect.DeepEqual(expected, actual) {
			t.Log(true)
		} else {
			t.Error("not equal!!")
			t.Error("actual= ", actual)
			t.Error("expected= ", expected)
		}
	}
}
