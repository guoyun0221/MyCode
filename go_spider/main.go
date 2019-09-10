package main

import (
	"fmt"
	"go_spider/spiders"
)

func main() {

	fmt.Println("Something you can do here (BTW, The code doesn't work for many webpages...)")
	fmt.Println("A : Get some emails from a page")
	fmt.Println("B : Get some pictures from a page")

	var s string
	fmt.Scanln(&s)

	if s == "A" || s == "a" {
		spiders.Email()
	} else if s == "B" || s == "b" {
		spiders.Img()
	}

}
