package p27_test

import (
	"reflect"
	"selfTesting/chall57/p27"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p27.Input
type TestOut = p27.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				FirstName:  "J",
				LastName:   "K",
				ZipCode:    "ABCDE",
				EmployeeID: "A12-1234",
			},
			output: TestOut{
				CheckFirstName:  false,
				CheckLastName:   false,
				CheckZipCode:    false,
				CheckEmployeeID: false,
			},
		},

		testData{
			input: TestIn{
				FirstName:  "Jimmy",
				LastName:   "James",
				ZipCode:    "55555",
				EmployeeID: "TK-4210",
			},
			output: TestOut{
				CheckFirstName:  true,
				CheckLastName:   true,
				CheckZipCode:    true,
				CheckEmployeeID: true,
			},
		},
	}

	return data
}

func TestPractice27Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p27.Practice27Normal(v.input)

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
