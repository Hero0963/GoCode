package p20_test

import (
	"reflect"
	"selfTesting/chall57/p20"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p20.Input
type TestOut = p20.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				Amount: 1000,
				State:  "Wisconsin",
				Town:   "Eau Claire",
			},
			output: TestOut{
				Tax:   5,
				Total: 1005,
			},
		},
	}

	return data
}

func TestPractice20Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p20.Practice20Normal(v.input)

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
