package p16_test

import (
	"reflect"
	"selfTesting/chall57/p16"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p16.Input
type TestOut = p16.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				Age: 16,
			},
			output: TestOut{
				Result: p16.ResultLegal,
			},
		},

		testData{
			input: TestIn{
				Age: 15,
			},
			output: TestOut{
				Result: p16.ResultNotLegal,
			},
		},

		testData{
			input: TestIn{
				Age: 0,
			},
			output: TestOut{
				Result: p16.ResultWrongInput,
			},
		},
	}

	return data
}

func TestPractice15Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p16.Practice16Normal(v.input)

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
