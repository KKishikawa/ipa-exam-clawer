package ipagojp

import (
	"clawer/modules/utilities"
	"strconv"
	"strings"
	"time"
)

// IPAの過去問題のURLをsliceで返す
func GetIPAExamUrls() []string {
	var now = time.Now()
	var startFiscalYear = now.Year()
	var currentMonth = int(now.Month())
	if currentMonth < 4 {
		startFiscalYear -= 1
	}
	// 2010年から今年までの年
	var years = utilities.Range(2010, startFiscalYear)
	var urls []string
	for _, year := range years {
		url := "https://www.ipa.go.jp/shiken/mondai-kaiotu/" + strconv.Itoa(year) + strings.ToLower(utilities.GetWareki(year, 4, true)) + ".html"
		urls = append(urls, url)
	}
	return urls
}
