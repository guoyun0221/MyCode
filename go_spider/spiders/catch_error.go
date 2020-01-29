package spiders

import "fmt"

func catch_error(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
