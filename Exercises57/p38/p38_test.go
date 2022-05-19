package p38_test

import (
	"reflect"
	"selfTesting/chall57/p38"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p38.Input
type TestOut = p38.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				RawData: "1 2 3 4 5 6 7 8",
			},
			output: TestOut{
				EvenNumbers: []int{2, 4, 6, 8},
			},
		},
	}

	return data
}

func TestPractice38Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p38.Practice38Normal(v.input)

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
