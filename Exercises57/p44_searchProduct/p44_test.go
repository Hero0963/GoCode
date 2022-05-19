package p44_test

import (
	"GoCode/Practices57/p44"
	"reflect"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = []p44.Input
type TestOut = p44.Output

func generateTestData() []testData {

	data := []testData{
		{
			input:  TestIn{},
			output: TestOut{},
		},
	}

	return data
}

func TestPractice44ormal(t *testing.T) {

	data := generateTestData()
	var name string

	for _, v := range data {

		var actual TestOut

		actual = p44.Practice44Normal(v.input, name)

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
