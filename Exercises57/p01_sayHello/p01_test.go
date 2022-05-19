package p01_test

import (
	"GoCode/Practices57/p01"
	"reflect"
	"testing"
)

type testData struct {
	name string
	ans  string
}

func generateTestData() []testData {
	data := []testData{
		testData{name: "Ted", ans: p01.StrPrefix + "Ted" + p01.StrSuffix},
		testData{name: "Hero", ans: p01.StrPrefix + "Hero" + p01.StrSuffix},
	}
	return data
}

func TestPractice01Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := p01.Practice01Normal(v.name)
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
