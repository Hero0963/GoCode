package p31_test

import (
	"reflect"
	"selfTesting/chall57/p31"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p31.Input
type TestOut = p31.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input:  TestIn{},
			output: TestOut{},
		},
	}

	return data
}

func TestPractice31Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p31.Practice31Normal(v.input)

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
