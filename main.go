package main

// 相対パスから絶対パスを取得するコマンドラインツール
// readline -fと同等の機能

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func main() {
	a := os.Args
	s := "/"
	if len(a) >= 2 {
		s = os.Args[1]
	} else {
		// 標準入力から読み込み
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		s = stdin.Text()
	}
	// 絶対パスへの変換
	p, err := filepath.Abs(s)
	if err != nil {
		log.Fatal(err)
	}
	// 出力
	println(p)
}
