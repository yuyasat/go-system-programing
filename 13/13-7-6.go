package main

import (
	"fmt"
	"sync"
)

func main() {
	smap := &sync.Map{}

	smap.Store("hello", "world")
	smap.Store(1, 2)

	smap.Delete("test")

	smap.Range(func(key, value interface{}) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})

	value, ok := smap.Load("hello")
	fmt.Printf("key=%v value=%v exists?=%v\n", "hello", value, ok)

	smap.LoadOrStore(1, 3)
	smap.LoadOrStore(2, 4)
	value1, ok1 := smap.Load(1)
	fmt.Printf("key=%v value=%v exists?=%v\n", 1, value1, ok1)

	value2, ok2 := smap.Load(2)
	fmt.Printf("key=%v value=%v exists?=%v\n", 1, value2, ok2)

}
