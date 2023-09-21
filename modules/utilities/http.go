package utilities

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// 指定されたURLのGoQueryオブジェクトを返す
func GetGoQueryFromUrl(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return goquery.NewDocumentFromReader(resp.Body)
}
