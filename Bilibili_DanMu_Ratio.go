package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type video struct {
	av    string
	Views int
	DanMu int
	ratio float64 // = danmu/views
}

const Sample_Size = 10000

func main() {
	fmt.Println("A:crawl B:analyze")
	var s string
	fmt.Scanln(&s)
	if s == "A" || s == "a" {
		crawl()
	} else if s == "B" || s == "b" {
		analyze()
	}
}

func crawl() {
	rand.Seed(time.Now().Unix())
	const Sample_Parts = 10
	const biggest_av = 75555555

	for cnt := 0; cnt < Sample_Size/Sample_Parts; { //cnt counts the number of valid videos//every time crawl a part
		url := "https://www.bilibili.com/video/av" + strconv.Itoa(rand.Intn(biggest_av))
		expr := `<a href="/video/(av[\d]+)[\S\s]+?(\d[\S\s]{0,10})播放[\S\s]+?(\d[\S\s]{0,10})弹幕`
		//get resp
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		//get body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		html := string(body)
		defer resp.Body.Close()
		//get match string list
		re := regexp.MustCompile(expr)
		match_list := re.FindAllStringSubmatch(html, -1)
		//save data
		for _, s := range match_list {
			Views := str2int(s[2])
			DanMu := str2int(s[3])
			if Views != 0 && DanMu != 0 { //don't count 0 views or 0 DanMu videos
				//open file
				f, err := os.OpenFile("videos_data.txt", os.O_CREATE|os.O_APPEND, 0666)
				if err != nil {
					fmt.Println(err)
				}
				defer f.Close()
				//write
				data := fmt.Sprintf("AV%sViews%sDanMu%s\n", s[1], s[2], s[3])
				_, err = f.WriteString(data)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Printf("%s  播放%s  弹幕%s\n", s[1], s[2], s[3]) //feedback
				cnt++                                            //valid data +1
			}
		}
	}
}

func analyze() {
	var videos []video
	var vid video
	//open file
	f, err := os.Open("videos_data.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	//turn string to struct slice
	buf := bufio.NewReader(f)
	for i := 0; i < Sample_Size; i++ {
		vid.av, vid.Views, vid.DanMu, vid.ratio = read_one_line(buf)
		videos = append(videos, vid)

	}
	//calculate
	var average_ratio, sum_ratio float64
	for i, vid := range videos {
		sum_ratio += vid.ratio
		fmt.Println(i, vid.ratio, sum_ratio)
	}
	average_ratio = sum_ratio / float64(len(videos))
	fmt.Printf("\nThe average ratio of DanMu to views is: %.3f%%\n", average_ratio*100)
	fmt.Println("sum ratio:", sum_ratio, "sample size:", len(videos))
}

func read_one_line(buf *bufio.Reader) (string, int, int, float64) {
	//get one line string
	line, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	//get index of av Views and DanMu
	index_AV := strings.Index(line, "AV")
	index_Views := strings.Index(line, "Views")
	index_DanMu := strings.Index(line, "DanMu")
	//transform them
	AV := line[index_AV+2 : index_Views]
	Views := str2int(line[index_Views+5 : index_DanMu])
	DanMu := str2int(line[index_DanMu+5 : len(line)-1]) //len(line)-1 is to get rid of '\n'
	Ratio := float64(DanMu) / float64(Views)

	return AV, Views, DanMu, Ratio
}

func str2int(s string) int {
	index1 := strings.Index(s, "万")
	index2 := strings.Index(s, ".")
	var ret int
	var err error

	if index1 == -1 { //doesn't reach 10k
		ret, err = strconv.Atoi(s) //just transform it
		if err != nil {
			fmt.Println(err)
		}
	} else { //reach 10k
		if index2 != -1 { //there is a decimal point
			part1, err := strconv.Atoi(s[:index2]) //the part of the left of the point
			if err != nil {
				fmt.Println(err)
			}
			part2, err := strconv.Atoi(s[index2+1 : index1]) //the right of the point
			if err != nil {
				fmt.Println(err)
			}
			ret = part1*10000 + part2*1000
		} else { //no decimal point
			raw_num, err := strconv.Atoi(s[:index1])
			if err != nil {
				fmt.Println(err)
			}
			ret = raw_num * 10000
		}
	}

	return ret
}
