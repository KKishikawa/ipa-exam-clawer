package ipagojp

// IPAの過去問題の区分の構造体
type IPAExamSubject struct {
	// 区分名
	name string
	// 問題のURL
	question_url string
	// 解答のURL
	answer_url string
	// 解説のURL
	comment_url *string
}

// IPAの過去問題の構造体
type IPAExam struct {
	// 年度
	year int
	// 1:春期 2:秋期
	season_type int
	// 試験名
	name string
	// 試験名(短縮形)
	short string
	// 試験問題区分
	subjects []IPAExamSubject
}
