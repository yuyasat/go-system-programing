package main

import (
	"fmt"
)

func main() {
	a := [4]int{1, 2, 3, 4}
	b := a[:]
	fmt.Println(&b[0], len(b), cap(b))

	c := a[1:3]
	fmt.Println(&c[0], len(c), cap(c))

	var d []int
	fmt.Println(len(d), cap(d))

	e := []int{1, 2, 3, 4}
	fmt.Println(&e[0], len(e), cap(e))

	f := make([]int, 4)
	fmt.Println(&f[0], len(f), cap(f))

	g := make([]int, 4, 8)
	fmt.Println(&g[0], len(g), cap(g))

	// 長さ1, 確保された要素2のスライスを作成
	s := make([]int, 1, 2)
	fmt.Println(&s[0], len(s), cap(s))

	// １要素追加（確保された範囲内）
	s = append(s, 1)
	fmt.Println(&s[0], len(s), cap(s))

	// さらに要素を追加（新しく配列が確保され直す）
	s = append(s, 2)
	fmt.Println(&s[0], len(s), cap(s)) // スライスの先頭要素のアドレスも変わる
}
