package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/test.txt")
	if err != nil { //err为error类型，常与空值表示
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name())
}
