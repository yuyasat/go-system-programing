package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var count int
	pool := sync.Pool{
		New: func() interface{} {
			count++
			return fmt.Sprintf("created: %d", count)
		},
	}

	pool.Put("manualy added: 1")
	pool.Put("manualy added: 2")
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get()) // これは新規作成

	pool.Put("removed 1")
	pool.Put("removed 2")
	pool.Put("removed 3")
	fmt.Println(pool.Get()) // removed 1
	runtime.GC()
	fmt.Println(pool.Get()) // 新規作成
}
