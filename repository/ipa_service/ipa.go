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

// 指定した年度のIPAの過去問題がDBに保存されているかを返す
func IsFullStoredIPAExam(year int) bool {
	db := utilities.Open()
	var count int64
	db.Model(&models.IPAExam{}).
		Where("year = ?", year).
		Count(&count)
	return count == 2
}
