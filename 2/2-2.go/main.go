package main

import (
	"fmt"
)

type Talker interface {
	Talk()
}

type Greeter struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Printf("Hello, my namge is %s\n", g.name)
}

func main() {
	var talker Talker
	talker = &Greeter{"wozozo"}
	talker.Talk()
}
