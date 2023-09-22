package main

import (
	"clawer/modules/ipa_go_jp"
	"fmt"

	flag "github.com/spf13/pflag"
)

func main() {
	flag.SetInterspersed(false)
	useHelp := flag.BoolP("help", "h", false, "ヘルプを表示します")
	useGetUrl := flag.BoolP("get-url", "g", false, "PDFのURLを取得します")
	flag.Parse()

	if *useHelp || flag.NFlag() == 0 {
		showHelp()
		return
	}

	if *useGetUrl {
		fmt.Println("PDFのURLを取得します")
		ipa_go_jp.Execute()
	}
}

// 実行時のコマンドライン引数のヘルプを表示します
func showHelp() {
	fmt.Println("このツールを利用して、IPAの過去問題の情報を取得できます。")
	flag.Usage()
}
