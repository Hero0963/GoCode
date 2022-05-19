package p25_test

import (
	"reflect"
	"selfTesting/chall57/p25"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p25.Input
type TestOut = p25.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				Password: "12345",
			},
			output: TestOut{
				Result: p25.ResultVeryWeak,
			},
		},

		testData{
			input: TestIn{
				Password: "abcdef",
			},
			output: TestOut{
				Result: p25.ResultWeak,
			},
		},

		testData{
			input: TestIn{
				Password: "abc123xyz",
			},
			output: TestOut{
				Result: p25.ResultStrong,
			},
		},

		testData{
			input: TestIn{
				Password: "1337h@xor",
			},
			output: TestOut{
				Result: p25.ResultVeryStrong,
			},
		},
	}

	return data
}

func TestPractice25Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p25.Practice25Normal(v.input)

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
