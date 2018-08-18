package main

import (
	"fmt"
	"time"
)

func sub() {
	fmt.Println("sub() is running")
	time.Sleep(time.Second)
	fmt.Println("sub() is finished")
}

func sub2() {
	fmt.Println("sub2() is running")
	time.Sleep(3 * time.Second)
	fmt.Println("sub2() is finished")
}

func main() {
	fmt.Println("start sub()")
	go sub2()
	time.Sleep(time.Second)
	go sub()
	time.Sleep(4 * time.Second)
}
