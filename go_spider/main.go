package main

import (
	"fmt"
	"go_spider/spiders"
)

func main() {

	fmt.Println("Something you can do here")
	fmt.Println(" ")
	fmt.Println("A: Get some emails from a page (BTW, The code doesn't work for many webpages...)")
	fmt.Println("B: Get some pictures from a page (The code doesn't work for many webpages, either...)")
	fmt.Println("C: Get 1000 random users' information of Bilibili, including id, name, sex, level and picture")

	var s string
	fmt.Scanln(&s)

	if s == "A" || s == "a" {
		spiders.Email()
	} else if s == "B" || s == "b" {
		spiders.Img()
	} else if s == "C" || s == "c" {
		spiders.Bilibili_user()
	}

}
