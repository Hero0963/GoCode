package p14_test

import (
	"reflect"
	"selfTesting/chall57/p14"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p14.Input
type TestOut = p14.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				Amount: 10,
				State:  "WI",
			},
			output: TestOut{
				SubTotal: 10,
				Tax:      0.55,
				Total:    10.55,
			},
		},

		testData{
			input: TestIn{
				Amount: 10,
				State:  "MN",
			},
			output: TestOut{
				SubTotal: 10,
				Tax:      0,
				Total:    10,
			},
		},
	}

	return data
}

func TestPractice14Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p14.Practice14Normal(v.input)

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
