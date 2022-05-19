package p34_test

import (
	"reflect"
	"selfTesting/chall57/p34"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p34.Input
type TestOut = p34.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input:  TestIn{},
			output: TestOut{},
		},
	}

	return data
}

func TestPractice34Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p34.Practice34Normal(v.input)

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
