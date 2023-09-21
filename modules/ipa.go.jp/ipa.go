package ipagojp

import (
	"clawer/modules/utilities"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	IPAStartYear = 2010
)

// IPAの過去問題のURLをsliceで返す
func GetIPAExamUrls() []string {
	var now = time.Now()
	var currentFiscalYear = now.Year()
	var currentMonth = int(now.Month())
	if currentMonth < 4 {
		currentFiscalYear -= 1
	}
	var urls []string
	for year := IPAStartYear; year < currentFiscalYear; year++ {
		url := "https://www.ipa.go.jp/shiken/mondai-kaiotu/" + strconv.Itoa(year) + strings.ToLower(utilities.GetWareki(year, 4, true)) + ".html"
		urls = append(urls, url)
	}
	return urls
}

// IPAの過去問題をdocumentから取得する
func GetIPAExamFromHTMLDoc(doc *goquery.Document) []IPAExam {
	var exams []IPAExam

	var titleText = doc.Find("title").Text()
	// タイトルから正規表現で年度を取得する
	year, err := strconv.Atoi(regexp.MustCompile(`\d{4}`).FindString(titleText))
	if err != nil {
		return exams
	}

	// 春期・秋期が文字列に含まれている要素を取得する
	var seasonTitleElements = doc.Find(".ttl")

	seasonTitleElements.Each(func(_ int, seasonTitleElement *goquery.Selection) {
		// タイトルから正規表現で春期・秋期を取得する
		var seasonTitleName = seasonTitleElement.Text()
		var seasonName = regexp.MustCompile(`.期`).FindString(seasonTitleName)
		var seasonType = 0
		if seasonName == "春期" {
			seasonType = 1
		} else if seasonName == "秋期" {
			seasonType = 2
		} else {
			return
		}

		// 試験名が文字列に含まれている要素を取得する
		var examTitleElements = seasonTitleElement.Next().Find(".anchorlink-list__item")
		examTitleElements.Each(func(_ int, examTitleElement *goquery.Selection) {
			// 試験名を取得する
			// 例：基本情報技術者試験（FE）
			var examTitle = examTitleElement.Text()
			// 例：基本情報技術者試験
			var examName = regexp.MustCompile(`（[A-Za-z]+）`).ReplaceAllString(examTitle, "")
			// 例：FE
			var examShort = regexp.MustCompile(`[A-Za-z]+`).FindString(examTitle)

			var exam = IPAExam{year, seasonType, examName, examShort, []IPAExamSubject{}}

			// 問題と試験区分が含まれる要素を取得する
			var examSubjectElements = doc.Find(examTitleElement.AttrOr("href", "") + " + * + * .def-list:not(:last-child)")
			examSubjectElements.Each(func(_ int, examSubjectElement *goquery.Selection) {
				// 試験区分名を取得する
				var examSubjectName = examSubjectElement.Find(".def-list__ttl").Text()

				var examSubject = IPAExamSubject{examSubjectName, "", "", ""}
				// 問題・解答・解説のURLを含む要素を取得する
				var examSubjectUrlElements = examSubjectElement.Find(".def-list__desc a")
				examSubjectUrlElements.Each(func(_ int, examSubjectUrlElement *goquery.Selection) {
					// URLを取得する
					var examSubjectUrl = examSubjectUrlElement.AttrOr("href", "")
					// URLを問題・解答・解説か判定する
					if strings.Contains(examSubjectUrl, "qs") {
						examSubject.question_url = examSubjectUrl
					} else if strings.Contains(examSubjectUrl, "ans") {
						examSubject.answer_url = examSubjectUrl
					} else if strings.Contains(examSubjectUrl, "cmnt") {
						examSubject.comment_url = examSubjectUrl
					}
				})
				exam.subjects = append(exam.subjects, examSubject)
			})

			exams = append(exams, exam)
		})
	})
	return exams

}
