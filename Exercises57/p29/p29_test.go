package p29_test

import (
	"reflect"
	"selfTesting/chall57/p29"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p29.Input
type TestOut = p29.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				Rate: "4",
			},
			output: TestOut{
				Year: 18,
			},
		},
	}

	return data
}

func TestPractice29Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p29.Practice29Normal(v.input)

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
