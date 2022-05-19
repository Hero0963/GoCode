package p24_test

import (
	"reflect"
	"selfTesting/chall57/p24"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p24.Input
type TestOut = p24.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				StrFisrt:  "note",
				StrSecond: "tone",
			},
			output: TestOut{
				Result: true,
			},
		},
	}

	return data
}

func TestPractice24Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p24.Practice24Normal(v.input)

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
