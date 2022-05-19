package p28_test

import (
	"reflect"
	"selfTesting/chall57/p28"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p28.Input
type TestOut = p28.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				Counter: 5,
				Numbers: []int{1, 2, 3, 4, 5},
			},
			output: TestOut{
				Total: 15,
			},
		},
	}

	return data
}

func TestPractice28Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p28.Practice28Normal(v.input)

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
