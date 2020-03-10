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
		vid = get_one_video(buf)                       //read a piece of info
		videos = insert(videos, 0, len(videos)-1, vid) //insert it to the slice in order by views
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

func insert(vids []Video, L, R int, vid Video) []Video {
	for L <= R {
		if vids[(L+R)/2].view < vid.view {
			L = (L+R)/2 + 1
		} else {
			R = (L+R)/2 - 1
		}
	}
	var t []Video
	t = append(t, vids[:L]...)
	t = append(t, vid)
	vids = append(t, vids[L:]...)
	return vids
}

func calculate(vids []Video) {
	var sum int
	var gradient [7]int
	//gradient[0]: view<10;gradient[1]: view>=10 && view<100;100-1k;1k-10k;10-100k;100k-1m;view>1m;
	for _, vid := range vids {
		sum += vid.view
		if vid.view < 10 {
			gradient[0]++
		} else if vid.view < 100 {
			gradient[1]++
		} else if vid.view < 1000 {
			gradient[2]++
		} else if vid.view < 10000 {
			gradient[3]++
		} else if vid.view < 100000 {
			gradient[4]++
		} else if vid.view < 1000000 {
			gradient[5]++
		} else {
			gradient[6]++
		}
	}
	fmt.Println("Sum:", sum)
	fmt.Println("Average:", sum/len(vids))
	fmt.Println("Median", vids[len(vids)/2].view)
	fmt.Println("---------Gradient table---------")
	fmt.Println("view<10", gradient[0])
	fmt.Println("10-100", gradient[1])
	fmt.Println("100~1k", gradient[2])
	fmt.Println("1k~10k", gradient[3])
	fmt.Println("10k~100k", gradient[4])
	fmt.Println("100k~1m", gradient[5])
	fmt.Println("view>1m", gradient[6])
}
