package main

import (
	"fmt"
	"go_spider/spiders"
)

func main() {
	fmt.Println(" ")
	fmt.Println("Something you can do here")
	fmt.Println("A: Get some emails from a page (BTW, The code doesn't work for many webpages...)")
	fmt.Println("B: Get some pictures from a page (The code doesn't work for many webpages, either...)")
	fmt.Println("C: Get 1000 random users' information of Bilibili, including id, name, sex, level and picture")
	fmt.Println("D: Get questions from zhihu(don't know how to pass zhihu's security verification, so it doesn't work sometimes)")
	fmt.Println("E: Get pictures from baidu")
	fmt.Println("F: Get top 250 books from Douban")

	var s string
	fmt.Scanln(&s)

	if s == "A" || s == "a" {
		spiders.Email()
	} else if s == "B" || s == "b" {
		spiders.Img()
	} else if s == "C" || s == "c" {
		spiders.Bilibili_user()
	} else if s == "D" || s == "d" {
		spiders.Zhihu_question()
	} else if s == "E" || s == "e" {
		spiders.Baidu_pic()
	} else if s == "F" || s == "f" {
		spiders.Douban_books()
	}

}
