package p18_test

import (
	"reflect"
	"selfTesting/chall57/p18"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p18.Input
type TestOut = p18.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				Unit:        "C",
				Temperature: 32,
			},
			output: TestOut{
				Result: 0,
			},
		},
	}

	return data
}

func TestPractice18Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p18.Practice18Normal(v.input)

		expected := v.output

		if reflect.DeepEqual(expected, actual) {
			t.Log(true)
		} else {
			t.Error("not equal!!")
			t.Error("actual= ", actual)
			t.Error("expected= ", expected)
		}
	}
}
