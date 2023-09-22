package ipa_go_jp

import (
	"clawer/modules/utilities"
	"clawer/repository/ipa_service"

	"fmt"
	"strconv"
)

// IPAの過去問題を取得し、DBに保存する
func Execute() {
	// IPAの過去問題の存在する年度を取得する
	var possiblyYears = getPossiblyIPAExamYears()
	// DBに保存されていない年度のみ過去問題の取得を行う
	for _, year := range possiblyYears {
		var storedSeasonTypes = ipa_service.GetSeasonTypesByYear(year)
		if len(storedSeasonTypes) == 2 {
			continue
		}

		// IPAの過去問題のURLを取得する
		var url = getIPAExamUrl(year)
		// documentを取得する
		doc, err := utilities.GetGoQueryFromUrl(url)
		if err != nil {
			fmt.Println(strconv.Itoa(year)+"年度のIPAの過去問題の取得に失敗しました", err)
			continue
		}
		// documentからIPAの過去問題を取得する
		var exams = getIPAExamFromHTMLDoc(doc, storedSeasonTypes...)
		// IPAの過去問題をDBに保存する
		ipa_service.SaveIPAExam(exams)
	}
}
