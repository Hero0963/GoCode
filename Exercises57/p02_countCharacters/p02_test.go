package p02_test

import (
	"GoCode/Practices57/p02"
	"reflect"
	"testing"
)

type testData struct {
	name string
	ans  string
}

func generateTestData() []testData {
	data := []testData{
		testData{name: "Ted", ans: "Ted" + p02.StrPrefix + "3" + p02.StrSuffix},
		testData{name: "Hero", ans: "Hero" + p02.StrPrefix + "4" + p02.StrSuffix},
	}
	return data
}

func TestPractice02Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := p02.Practice02Normal(v.name)
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
