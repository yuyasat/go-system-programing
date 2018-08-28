package main

import (
	"fmt"
	"sync"
)

func initialize() {
	fmt.Println("初期化処理")
}

var once sync.Once

func main() {
	once.Do(initialize)
	once.Do(initialize)
	once.Do(initialize)
}
