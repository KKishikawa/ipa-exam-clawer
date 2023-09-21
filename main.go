package main

import (
	"clawer/modules/utilities"
	"fmt"
)

func main() {
	const (
		sampleUrl = "https://golang.org"
	)
	// HttpGet()を呼び出す
	body, err := utilities.HttpGet(sampleUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
}
