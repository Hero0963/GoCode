package p40

import (
	"sort"
	"strings"
)

func GetData() []Input {

	data := []Input{
		{
			FirstName:      "John",
			LastName:       "Johnson",
			Position:       "Manager",
			SeparationDate: "2016-12-31",
		},

		{
			FirstName:      "Tou",
			LastName:       "Xiong",
			Position:       "Software Engineer",
			SeparationDate: "2016-10-05",
		},

		{
			FirstName:      "Michaela",
			LastName:       "Michaelson",
			Position:       "District Manager",
			SeparationDate: "2015-12-19",
		},

		{
			FirstName:      "Jake",
			LastName:       "Jacobson",
			Position:       "Programmer",
			SeparationDate: "2021-02-01",
		},

		{
			FirstName:      "Jacquelyn",
			LastName:       "Jackson",
			Position:       "DBA",
			SeparationDate: "2021-02-01",
		},

		{
			FirstName:      "Sally",
			LastName:       "Webber",
			Position:       "Web Developer",
			SeparationDate: "2015-12-18",
		},
	}

	return data
}

func filterEmployeeListByName(input []Input, s string) []Output {

	output := make([]Output, 0, len(input))

	for _, val := range input {

		var fullName string
		fullName = val.FirstName + " " + val.LastName

		if strings.Contains(fullName, s) {
			var oput Output
			oput.FullName = fullName
			oput.Position = val.Position
			oput.SeparationDate = val.SeparationDate

			output = append(output, oput)

		}

	}

	sort.Sort(ByName(output))

	return output
}
