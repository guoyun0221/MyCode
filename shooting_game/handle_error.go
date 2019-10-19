package main

import "fmt"

func handle_error(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
