package ipa_go_jp

import (
	"clawer/models"
	"clawer/modules/utilities"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	startYearOfExam = 2010
	domainIPA       = "https://www.ipa.go.jp"
	mondaiKaitouUrl = domainIPA + "/shiken/mondai-kaiotu/"
)

// IPAの過去問題の存在する年度をsliceで返す
func getPossiblyIPAExamYears() []int {
	var now = time.Now()
	var currentFiscalYear = now.Year()
	var currentMonth = int(now.Month())
	if currentMonth < 4 {
		currentFiscalYear -= 1
	}

	var years = make([]int, currentFiscalYear-startYearOfExam+1)
	for i := range years {
		years[i] = startYearOfExam + i
	}
	return years
}

// IPAの過去問題のURLをsliceで返す
func getIPAExamUrl(year int) string {
	var gengoInfo = utilities.GetWareki(year, 4)
	var gengoStr = strings.ToLower(gengoInfo.Era) + fmt.Sprintf("%02d", gengoInfo.Year)
	return mondaiKaitouUrl + strconv.Itoa(year) + gengoStr + ".html"
}

// IPAの過去問題をdocumentから取得する
func getIPAExamFromHTMLDoc(doc *goquery.Document, excludeSeasonTypes ...uint8) []*models.IPAExam {
	var exams []*models.IPAExam

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
		var seasonType uint8 = 0
		if seasonName == "春期" {
			seasonType = 1
		} else if seasonName == "秋期" {
			seasonType = 2
		} else {
			return
		}
		// excludeSeasonTypesに含まれている場合は処理をスキップする
		if utilities.Contains(&excludeSeasonTypes, seasonType) {
			return
		}
		var exam = &models.IPAExam{Year: year, SeasonType: seasonType, ExamTypes: []*models.IPAExamType{}}
		exams = append(exams, exam)

		// 試験区分名が文字列に含まれている要素を取得する
		var examTitleElements = seasonTitleElement.Next().Find(".anchorlink-list__item")
		examTitleElements.Each(func(_ int, examTitleElement *goquery.Selection) {
			// 試験区分名を取得する
			// 例：基本情報技術者試験（FE）
			var examTitle = examTitleElement.Text()
			// 例：基本情報技術者試験
			var examName = regexp.MustCompile(`（[A-Za-z]+）`).ReplaceAllString(examTitle, "")
			// 例：FE
			var examShort = regexp.MustCompile(`[A-Za-z]+`).FindString(examTitle)

			var examType = &models.IPAExamType{Name: examName, Short: examShort, Subjects: []*models.IPAExamSubject{}}
			exam.ExamTypes = append(exam.ExamTypes, examType)

			// 問題と試験区分が含まれる要素を取得する
			var examSubjectElements = doc.Find(examTitleElement.AttrOr("href", "") + " + * + * .def-list:not(:last-child)")
			examSubjectElements.Each(func(_ int, examSubjectElement *goquery.Selection) {
				// 試験区分名を取得する
				var examSubjectName = examSubjectElement.Find(".def-list__ttl").Text()

				var examSubject = &models.IPAExamSubject{Name: examSubjectName, IPAExamData: []*models.IPAExamData{}}
				examType.Subjects = append(examType.Subjects, examSubject)

				// 問題・解答・解説のURLを含む要素を取得する
				var examSubjectUrlElements = examSubjectElement.Find(".def-list__desc a")
				examSubjectUrlElements.Each(func(_ int, examSubjectUrlElement *goquery.Selection) {
					// URLを取得する
					var examSubjectUrl = examSubjectUrlElement.AttrOr("href", "")
					// urlが/で始まっている場合はドメインを付与する
					if strings.HasPrefix(examSubjectUrl, "/") {
						examSubjectUrl = domainIPA + examSubjectUrl
					}
					// URLを問題・解答・解説か判定する
					var dataType uint8 = 0
					if strings.Contains(examSubjectUrl, "qs") {
						dataType = 1
					} else if strings.Contains(examSubjectUrl, "ans") {
						dataType = 2
					} else if strings.Contains(examSubjectUrl, "cmnt") {
						dataType = 3
					} else {
						return
					}
					var examSubjectData = &models.IPAExamData{Url: examSubjectUrl, DataType: dataType}
					examSubject.IPAExamData = append(examSubject.IPAExamData, examSubjectData)
				})
			})
		})
	})
	return exams

}
