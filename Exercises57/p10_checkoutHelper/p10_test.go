package p10_test

import (
	"GoCode/Practices57/p10"
	"reflect"

	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = []p10.Item
type TestOut = p10.CalPrice

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				p10.Item{Price: 25, Quantity: 2},
				p10.Item{Price: 10, Quantity: 1},
				p10.Item{Price: 4, Quantity: 1},
			},
			output: TestOut{
				SubTotal: 64, Tax: 3.52, Total: 67.52,
			},
		},
	}

	return data
}

func TestPractice10Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p10.Practice10Normal(v.input)

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
