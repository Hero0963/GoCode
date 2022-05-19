package p04_test

import (
	"GoCode/Practices57/p04"
	"reflect"
	"testing"
)

type testData struct {
	noun string
	verb string
	adj  string
	adv  string
	ans  string
}

func generateTestData() []testData {
	data := []testData{
		testData{noun: "dog", verb: "walk", adj: "blue", adv: "quickly", ans: p04.SentancePrefix1 + "walk" + p04.SentancePrefix2 + "blue" + " " + "dog" + " " + "quickly" + p04.SentancePrefix3},
	}
	return data
}

func TestPractice04Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := p04.Practice04Normal(v.noun, v.verb, v.adj, v.adv)
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
