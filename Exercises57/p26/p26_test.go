package p26_test

import (
	"reflect"
	"selfTesting/chall57/p26"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p26.Input
type TestOut = p26.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				Balance:    5000,
				APR:        12,
				MonthlyPay: 100,
			},
			output: TestOut{
				NeedMonth: 70,
			},
		},
	}

	return data
}

func TestPractice26Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p26.Practice26Normal(v.input)

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
