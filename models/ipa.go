package models

import (
	"gorm.io/gorm"
)

// IPAの過去問題の試験科目の構造体
type IPAExamSubject struct {
	gorm.Model
	// 試験科目名
	Name string
	// 問題のURL
	QuestionUrl string
	// 解答のURL
	AnswerUrl string
	// 解説のURL
	CommentUrl string
	// 試験科目のID
	IPAExamTypeID uint
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
	SeasonType int `gorm:"uniqueIndex:idx_year_season_type"`
	// 試験区分
	ExamTypes []*IPAExamType
}
