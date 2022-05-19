package p06_test

import (
	"GoCode/Practices57/p06"
	"reflect"
	"testing"
)

type testData struct {
	curAge    int
	retireAge int

	ans testAnswer
}

type testAnswer struct {
	diff int
}

func generateTestData() []testData {
	data := []testData{
		testData{curAge: 25, retireAge: 65, ans: testAnswer{diff: 40}},
		testData{curAge: -1, retireAge: 65, ans: testAnswer{diff: 0}},
		testData{curAge: 35, retireAge: 25, ans: testAnswer{diff: -10}},
	}
	return data
}

func TestPractice06Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actualDiff := p06.Practice06Normal(v.curAge, v.retireAge)

		var actual testAnswer
		actual.diff = actualDiff

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
