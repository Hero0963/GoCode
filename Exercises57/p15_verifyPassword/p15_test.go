package p15_test

import (
	"reflect"
	"selfTesting/chall57/p15"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p15.Input
type TestOut = p15.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				UserID:   "hero",
				Password: "abc$123",
			},
			output: TestOut{
				Result: p15.ResultMatch,
			},
		},

		testData{
			input: TestIn{
				UserID:   "ted",
				Password: "456789",
			},
			output: TestOut{
				Result: p15.ResultNoUserID,
			},
		},

		testData{
			input: TestIn{
				UserID:   "hero",
				Password: "yerty",
			},
			output: TestOut{
				Result: p15.ResultWrongPWD,
			},
		},
	}

	return data
}

func TestPractice15Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p15.Practice15Normal(v.input)

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
