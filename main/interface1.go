package main

import (
	"fmt"
)

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}
func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

func main() {
	fmt.Println("hello")
}
