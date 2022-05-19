package p07_test

import (
	"GoCode/Practices57/p07"
	"reflect"
	"testing"
)

type testData struct {
	areaSquareFeet float64
	ans            testAnswer
}

type testAnswer struct {
	areaSquareMeters float64
}

func generateTestData() []testData {
	data := []testData{
		testData{areaSquareFeet: 300, ans: testAnswer{areaSquareMeters: 27.870912}},
	}
	return data
}

func TestPractice07Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		tmp := p07.Practice07Normal(v.areaSquareFeet)

		var actual testAnswer
		actual.areaSquareMeters = tmp

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
