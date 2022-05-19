package p11_test

import (
	"GoCode/Practices57/p11"
	"reflect"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p11.Input
type TestOut = p11.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				Euros:        81,
				ExchangeRate: 137.51,
			},
			output: TestOut{
				USDollars: 111.38,
			},
		},
	}

	return data
}

func TestPractice10Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p11.Practice11Normal(v.input)

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
