package utilities

import (
	"io"
	"net/http"
)

// HttpGet()は、指定されたURLのHTMLを取得する
func HttpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// HTMLを取得する
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
