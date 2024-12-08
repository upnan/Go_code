package main

import (
	"fmt"
)

type author struct { //作者结构体
	firstName string
	lastName  string
	boi       string
}

func (a author) fullname() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

//博客文结构体

type blogPost struct {
	title   string
	content string
	author  //匿名字段代表一个作者
}

func (b blogPost) details() {
	fmt.Println("Title:", b.title)
	fmt.Println("Content:", b.content)
	/*每当一个结构体字段嵌套在另一个结构体中时，Go 语言提供了将嵌套的字段视为外层结构体一部分的选项。
	这意味着上述代码第 11 行中的 p.author.fullName() 可以被替换为 p.fullName() 。*/
	// fmt.Println("Auther:",b.author.fullname())
	// fmt.Println("Bio:",b.author.boi)
	fmt.Println("Auther:", b.fullname())
	fmt.Println("Bio:", b.boi)
}

// 创建网站结构体
type website struct {
	blogPosts []blogPost //一个blogpost类型的切片
}

func (w website) contents() {
	fmt.Println("Contents of Website")
	fmt.Println()
	for _, v := range w.blogPosts {
		v.details()
		fmt.Println()
	}
}

func main() {
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	blogPost1 := blogPost{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	blogPost2 := blogPost{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	blogPost3 := blogPost{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	w := website{
		blogPosts: []blogPost{blogPost1, blogPost2, blogPost3},
	}
	w.contents()
}
