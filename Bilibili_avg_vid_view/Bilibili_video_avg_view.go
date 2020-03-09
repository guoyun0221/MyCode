package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Sample_Size = 10000

type Video struct {
	av, view, like, coin, favorite, danmaku, reply int
}

func main() {
	videos := make([]Video, 0, Sample_Size)
	var vid Video

	f, err := os.Open("videos_info.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	buf := bufio.NewReader(f)

	for i := 0; i < Sample_Size; i++ {
		vid = get_one_video(buf)     //read a piece of info
		videos = insert(videos, vid) //insert it to the slice in order by views
	}
	calculate(videos)
}

func get_one_video(buf *bufio.Reader) Video {
	line, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	var vid Video
	//find the index of substrs
	title_index := strings.Index(line, ", 标题:")
	view_index := strings.Index(line, ", 播放:")
	danmaku_index := strings.Index(line, ", 弹幕:")
	// str to int , give it to vid
	vid.av, _ = strconv.Atoi(line[3:title_index])
	vid.view, _ = strconv.Atoi(line[view_index+9 : danmaku_index])

	return vid
}

func insert(vids []Video, vid Video) []Video {
	var i int
	for i = 0; i < len(vids); i++ {
		if vids[i].view >= vid.view {
			break
		}
	}
	var t []Video
	t = append(t, vids[:i]...)
	t = append(t, vid)
	vids = append(t, vids[i:]...)
	return vids
}

func calculate(vids []Video) {
	var sum int
	var gradient [11]int
	//gradient[0]: view<50;gradient[1]: view>=50 && view<100;100-500;500-1k;1k-5k;5k-10k;10k-50k;
	//50k-100k;100k-500k;500k-1m;view>1m;
	for _, vid := range vids {
		sum += vid.view
		if vid.view < 50 {
			gradient[0]++
		} else if vid.view < 100 {
			gradient[1]++
		} else if vid.view < 500 {
			gradient[2]++
		} else if vid.view < 1000 {
			gradient[3]++
		} else if vid.view < 5000 {
			gradient[4]++
		} else if vid.view < 10000 {
			gradient[5]++
		} else if vid.view < 50000 {
			gradient[6]++
		} else if vid.view < 100000 {
			gradient[7]++
		} else if vid.view < 500000 {
			gradient[8]++
		} else if vid.view < 1000000 {
			gradient[9]++
		} else {
			gradient[10]++
		}
	}
	fmt.Println("Sum:", sum)
	fmt.Println("Average:", sum/len(vids))
	fmt.Println("Median", vids[len(vids)/2].view)
	fmt.Println("---------Gradient table---------")
	fmt.Println("view<50", gradient[0])
	fmt.Println("50-100", gradient[1])
	fmt.Println("100~500", gradient[2])
	fmt.Println("500~1k", gradient[3])
	fmt.Println("1k~5k", gradient[4])
	fmt.Println("5k~10k", gradient[5])
	fmt.Println("10k~50k", gradient[6])
	fmt.Println("50k~100k", gradient[7])
	fmt.Println("100k~500k", gradient[8])
	fmt.Println("500k~1m", gradient[9])
	fmt.Println("view>1m", gradient[10])
}
