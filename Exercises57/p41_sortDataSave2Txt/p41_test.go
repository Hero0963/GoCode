package p41_test

import (
	"GoCode/Practices57/p41"
	"reflect"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p41.Input
type TestOut = p41.Output

func generateTestData() []testData {

	data := []testData{
		{
			input: TestIn{

				RawData: []string{"Ling, Mai",
					"Johnson, Jim",
					"Zarnecki, Sabrina",
					"Jones, Chris",
					"Jones, Aaron",
					"Swift, Geoffrey",
					"Xiong, Fong"},
			},
			output: TestOut{

				SortedData: []string{"Johnson, Jim",
					"Jones, Aaron",
					"Jones, Chris",
					"Ling, Mai",
					"Swift, Geoffrey",
					"Xiong, Fong",
					"Zarnecki, Sabrina"},
			},
		},
	}

	return data
}

func TestPractice41ormal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p41.Practice41Normal(v.input)

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
