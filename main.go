package main

import (
	ipagojp "clawer/modules/ipa.go.jp"
	"clawer/modules/utilities"
	"fmt"
)

func main() {
	url := ipagojp.GetIPAExamUrls()[0]
	// HttpGet()を呼び出す
	doc, err := utilities.GetGoQueryFromUrl(url)
	if err != nil {
		panic(err)
	}
	exams := ipagojp.GetIPAExamFromHTMLDoc(doc)
	fmt.Printf("%+v\n", exams)
}
