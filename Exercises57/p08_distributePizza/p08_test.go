package p08_test

import (
	"GoCode/Practices57/p08"
	"reflect"
	"testing"
)

type testData struct {
	numberOfPeople int
	numberOfPizza  int
	ans            testAnswer
}

type testAnswer struct {
	numberOfShared   int
	numberOfLeftover int
}

func generateTestData() []testData {
	data := []testData{
		testData{numberOfPeople: 8, numberOfPizza: 2, ans: testAnswer{numberOfShared: 2, numberOfLeftover: 0}},
	}
	return data
}

func TestPractice08Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual testAnswer

		actual.numberOfShared, actual.numberOfLeftover = p08.Practice08Normal(v.numberOfPeople, v.numberOfPizza)

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
