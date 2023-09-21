package utilities

import (
	"strconv"
)

type Wareki struct {
	year  int
	month int
	name  string
	era   string
}

func GetWareki(year int, month int, isRyaku bool) string {
	warekiSettings := []Wareki{
		{year: 2019, month: 5, name: "令和", era: "R"},
		{year: 1989, month: 1, name: "平成", era: "H"},
		{year: 1926, month: 12, name: "昭和", era: "S"},
		{year: 1912, month: 7, name: "大正", era: "T"},
		{year: 1868, month: 1, name: "明治", era: "M"},
	}

	for _, w := range warekiSettings {
		if year > w.year || (year == w.year && month >= w.month) {
			if isRyaku {
				return w.era + strconv.Itoa(year-w.year+1)
			} else {
				var warekiY = year - w.year + 1
				var warekiStr string
				if warekiY == 1 {
					warekiStr = "元"
				} else {
					warekiStr = strconv.Itoa(warekiY)
				}
				return w.name + warekiStr + "年"
			}
		}
	}
	return ""
}
