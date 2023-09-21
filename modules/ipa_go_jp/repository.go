package ipa_go_jp

import (
	"clawer/models"
	"clawer/modules/utilities"

	"gorm.io/gorm"
)

// 取得したIPAの過去問題をDBに保存する
func saveIPAExam(exams []*models.IPAExam) {
	db := utilities.Open()
	db.Transaction(func(tx *gorm.DB) error {
		for _, exam := range exams {
			tx.Create(exam)
		}
		return nil
	})
}

// 指定した年度のIPAの過去問題がDBに保存されているかを返す
func isStoredIPAExam(year int) bool {
	db := utilities.Open()
	var exists bool
	db.Model(&models.IPAExam{}).
		Where("year = ?", year).
		Limit(1).
		Select("1").
		Find(&exists)
	return exists
}
