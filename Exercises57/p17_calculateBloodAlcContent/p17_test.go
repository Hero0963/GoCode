package p17_test

import (
	"reflect"
	"selfTesting/chall57/p17"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p17.Input
type TestOut = p17.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				Amount: 20,
				Weight: 90,
				Hour:   2,
				Sex:    1,
			},
			output: TestOut{
				BAC:    1.53,
				Result: p17.ResultNotLegal,
			},
		},
	}

	return data
}

func TestPractice17Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p17.Practice17Normal(v.input)

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
