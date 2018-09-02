package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(5 * time.Second)
	fmt.Println(10 * time.Millisecond)
	fmt.Println(time.ParseDuration("10m30s"))

	// 現在時刻
	time.Now()
	// 指定日時を作成
	fmt.Println(time.Date(2017, time.August, 26, 11, 50, 30, 0, time.Local))
	// フォーマットを指定してパース
	fmt.Println(time.Parse(time.Kitchen, "11:30AM"))
	// Epochタイムから作成
	fmt.Println(time.Unix(1503673200, 0))

	// 3時間後の時間
	fmt.Println(time.Now().Add(3 * time.Hour))
	// ファイルの変更日時が何日まえか知る
	fileInfo, _ := os.Stat("/Users/Yuya/.vimrc")
	fmt.Printf("%v前 \n", time.Now().Sub(fileInfo.ModTime()))
	// 時間を１時間ごとに丸める
	fmt.Println(time.Now().Round(time.Hour))
}
