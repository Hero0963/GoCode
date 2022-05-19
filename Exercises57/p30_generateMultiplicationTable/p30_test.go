package p30_test

import (
	"reflect"
	"selfTesting/chall57/p30"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p30.Input
type TestOut = p30.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input:  TestIn{},
			output: TestOut{},
		},
	}

	return data
}

func TestPractice30Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p30.Practice30Normal(v.input)

		expected := v.output

		if reflect.DeepEqual(expected, actual) {
			t.Log(true)
		} else {
			t.Error("not equal!!")
			t.Error("input= ", v.input)
			t.Error("actual= ", actual)
			t.Error("expected= ", expected)
		}
	}
}
