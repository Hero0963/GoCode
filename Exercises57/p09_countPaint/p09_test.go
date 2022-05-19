package p09_test

import (
	"GoCode/Practices57/p09"
	"reflect"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn struct {
	areaCeiling int
}

type TestOut struct {
	numberOfGallon int
}

func generateTestData() []testData {
	data := []testData{
		testData{input: TestIn{areaCeiling: 360}, output: TestOut{numberOfGallon: 2}},
	}
	return data
}

func TestPractice09Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual.numberOfGallon = p09.Practice09Normal(v.input.areaCeiling)

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
