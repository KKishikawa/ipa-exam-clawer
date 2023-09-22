package ipa_service

import (
	"clawer/models"
	"clawer/modules/utilities"

	"gorm.io/gorm"
)

// 取得したIPAの過去問題をDBに保存する
func SaveIPAExam(exams []*models.IPAExam) {
	db := utilities.Open()
	db.Transaction(func(tx *gorm.DB) error {
		for _, exam := range exams {
			tx.Create(exam)
		}
		return nil
	})
}

// 指定した年度のIPAの過去問題のうち登録済みの時期を返す
func GetSeasonTypesByYear(year int) []uint8 {
	db := utilities.Open()
	var seasonTypes []uint8
	db.Model(&models.IPAExam{}).
		Where("year = ?", year).
		Pluck("season_type", &seasonTypes)
	return seasonTypes
}
