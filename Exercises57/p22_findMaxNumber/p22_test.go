package p22_test

import (
	"reflect"
	"selfTesting/chall57/p22"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p22.Input
type TestOut = p22.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				Numbers: []int{4, 90, 111, 66},
			},
			output: TestOut{
				Numbers: []int{4, 66, 90, 111},
			},
		},
	}

	return data
}

func TestPractice22Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p22.Practice22Normal(v.input)

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
