package p21_test

import (
	"reflect"
	"selfTesting/chall57/p21"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p21.Input
type TestOut = p21.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				Number: 4,
			},
			output: TestOut{
				Name: "April",
			},
		},
	}

	return data
}

func TestPractice21Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p21.Practice21Normal(v.input)

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
