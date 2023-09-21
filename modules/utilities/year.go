package utilities

type warekiConf struct {
	year  int
	month int
	name  string
	era   string
}

type GengoInfo struct {
	Name string
	Era  string
	Year int
}

func GetWareki(year int, month int) GengoInfo {
	warekiSettings := []warekiConf{
		{year: 2019, month: 5, name: "令和", era: "R"},
		{year: 1989, month: 1, name: "平成", era: "H"},
		{year: 1926, month: 12, name: "昭和", era: "S"},
		{year: 1912, month: 7, name: "大正", era: "T"},
		{year: 1868, month: 1, name: "明治", era: "M"},
	}

	for _, w := range warekiSettings {
		if year > w.year || (year == w.year && month >= w.month) {
			return GengoInfo{
				Name: w.name,
				Era:  w.era,
				Year: year - w.year + 1,
			}
		}
	}
	return GengoInfo{}
}
