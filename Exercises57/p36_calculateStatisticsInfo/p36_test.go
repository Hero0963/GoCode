package p36_test

import (
	"reflect"
	"selfTesting/chall57/p36"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p36.Input
type TestOut = p36.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				Numbers: []float64{100, 200, 1000, 300},
			},
			output: TestOut{
				Mean: 400,
				Max:  1000,
				Min:  100,
				StdD: 353.55,
			},
		},
	}

	return data
}

func TestPractice36Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p36.Practice36Normal(v.input)

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
