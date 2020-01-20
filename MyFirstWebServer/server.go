package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":5555", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("webpage.html")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	_, err = io.Copy(w, f)
	if err != nil {
		fmt.Println(err)
	}
}
