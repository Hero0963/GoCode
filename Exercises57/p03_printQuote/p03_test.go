package p03_test

import (
	"GoCode/Practices57/p03"
	"reflect"
	"testing"
)

type testData struct {
	s1  string
	s2  string
	ans string
}

func generateTestData() []testData {
	data := []testData{
		testData{s1: "These aren't the droids you're looking for.", s2: "Obi-Wan Kenobi", ans: "Obi-Wan Kenobi" + p03.StrPrefix + "These aren't the droids you're looking for."},
	}
	return data
}

func TestPractice03Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := p03.Practice03Normal(v.s1, v.s2)
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
