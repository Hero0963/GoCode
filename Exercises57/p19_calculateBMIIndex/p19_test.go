package p19_test

import (
	"reflect"
	"selfTesting/chall57/p19"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p19.Input
type TestOut = p19.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				Height: 100,
				Weight: 100,
			},
			output: TestOut{
				BMI:    7.03,
				Result: "UnderWeight",
			},
		},
	}

	return data
}

func TestPractice18Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p19.Practice19Normal(v.input)

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
