package models

import (
	"time"

	"gorm.io/gorm"
)

type IPAExamData struct {
	// IPAの過去問題の試験科目のID
	IPAExamSubjectId uint `gorm:"primaryKey"`
	// データの種類　1:問題 2:解答 3:解説
	DataType uint8 `gorm:"primaryKey"`
	// データのURL
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// IPAの過去問題の試験科目の構造体
type IPAExamSubject struct {
	gorm.Model
	// 試験科目名
	Name string
	// 試験科目のID
	IPAExamTypeID uint
	// 試験科目のデータ
	IPAExamData []*IPAExamData
}

// IPAの過去問題の試験区分の構造体
type IPAExamType struct {
	gorm.Model
	// 試験区分名
	Name string
	// 試験区分名(短縮形)
	Short string
	// IPAの過去問題のID
	IPAExamID uint
	// 試験科目
	Subjects []*IPAExamSubject
}

// IPAの過去問題の構造体
type IPAExam struct {
	gorm.Model
	// 年度
	Year int `gorm:"uniqueIndex:idx_year_season_type"`
	// 1:春期 2:秋期
	SeasonType uint8 `gorm:"uniqueIndex:idx_year_season_type"`
	// 試験区分
	ExamTypes []*IPAExamType
}
