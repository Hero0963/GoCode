package p39

import "sort"

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

func sortEmployeeListByLastName(input []Input) []Output {

	output := make([]Output, 0, len(input))

	m := make(map[string]Input)

	for _, v := range input {
		m[v.LastName] = v
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		val := m[k]

		var oput Output
		oput.FullName = val.FirstName + " " + val.LastName
		oput.Position = val.Position
		oput.SeparationDate = val.SeparationDate

		output = append(output, oput)

	}

	return output
}
