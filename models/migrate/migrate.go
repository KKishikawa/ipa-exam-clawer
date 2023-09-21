package migrate

import (
	"clawer/models"

	"gorm.io/gorm"
)

func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(
		&models.IPAExam{},
		&models.IPAExamType{},
		&models.IPAExamSubject{},
	)
}
