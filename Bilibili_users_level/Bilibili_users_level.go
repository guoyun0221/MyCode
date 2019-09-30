package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var data [7]int // index is level, from 0 to 6; value is the number of users in this level
	var sample_size int = 10000
	var progress int //to show progress of the work by print star

	file, err := os.Open("Bilibili_users.txt") //open file of users' info //the file here is a point to the file
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()           //file is also a io.reader
	buf := bufio.NewReader(file) //buf is a *Reader, I don't really know what it is.

	for i := 0; i < sample_size; i++ { //read users' info and count
		lv := read_one_user(buf)
		data[lv]++
		print_progress(&progress, sample_size)
		progress++
		time.Sleep(time.Second / 1000) // act like there is huge number of data :) // or it will suddenly and rapidly to the end
	}

	fmt.Println("")
	show_result(data, sample_size)
}

func read_one_user(buf *bufio.Reader) int {

	sub := "lv:"

	line, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	index := strings.Index(line, sub)
	lv, err := strconv.Atoi(line[index+3 : index+4]) //trans lv in string to int

	return lv
}

func print_progress(num *int, size int) *int {
	step := size / 100 //print one star every 1 percent
	if *num >= step {
		fmt.Printf("*")
		*num -= step
	}
	return num
}

func show_result(data [7]int, size int) {
	for i, _ := range data {
		percent := float32(data[i]) / float32(size) * 100.0
		fmt.Printf("lv:%d \t number:%d \t proportion:%.2f%% \n", i, data[i], percent)
	}
}
