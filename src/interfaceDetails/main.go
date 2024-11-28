package main

import (
	"fmt"
)

type Interface interface {
	say()
}

type stu struct {
	Name string
}

func (s stu) say() {
	fmt.Println("say")
}

func main() {
	var Stu stu
	var a Interface = Stu
	a.say()
}
