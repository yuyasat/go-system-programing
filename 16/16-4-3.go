package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("waiting 5 seconds")
	after := time.After(5 * time.Second)
	<-after
	fmt.Println("done")

	fmt.Println("waiting 5 seconds")
	for now := range time.Tick(5 * time.Second) {
		fmt.Println("now:", now)
	}
}
