package spiders

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Resp struct {
	Data Video `json:"data"`
}

type Video struct {
	Aid   int    `json:"aid"`
	Title string `json:"title"`
	Owner Up     `json:"owner"`
	Stat  State  `json:"stat"`
}

type Up struct {
	Name string `json:"name"`
}

type State struct {
	View     int `json:"view"`
	Danmaku  int `json:"danmaku"`
	Reply    int `json:"reply"`
	Favorite int `json:"favorite"`
	Coin     int `json:"coin"`
	Share    int `json:"share"`
	Like     int `json:"like"`
}

func Bilibili_videos_info() {
	Max_aid := 93436487 //current max av number
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000; i++ {
		aid := rand.Intn(Max_aid)
		url := "https://api.bilibili.com/x/web-interface/view?aid=" + strconv.Itoa(aid)
		get_video_info(url)
		time.Sleep(time.Second / 10)
	}

}

func get_video_info(url string) {
	client := &http.Client{} //These below is to avoid 403
	req, err := http.NewRequest("GET", url, nil)
	catch_error(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:69.0) Gecko/20100101 Firefox/69.0")
	res, err := client.Do(req)
	catch_error(err)
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var resp Resp
	err = json.Unmarshal(body, &resp)
	catch_error(err)

	if resp.Data.Aid != 0 {
		save_and_output(resp)
	}

}

func save_and_output(r Resp) {
	fmt.Printf("av:%d, 标题:%s, up:%s, 播放:%d, 弹幕:%d, 点赞:%d, 评论%d, 投币%d, 收藏%d,\n", r.Data.Aid, r.Data.Title, r.Data.Owner.Name, r.Data.Stat.View, r.Data.Stat.Danmaku, r.Data.Stat.Like, r.Data.Stat.Reply, r.Data.Stat.Coin, r.Data.Stat.Favorite)
	f, err := os.OpenFile("videos_info.txt", os.O_CREATE|os.O_APPEND, 0666)
	catch_error(err)
	defer f.Close()
	info := fmt.Sprintf("av:%d, 标题:%s, up:%s, 播放:%d, 弹幕:%d, 点赞:%d, 评论%d, 投币%d, 收藏%d,\n", r.Data.Aid, r.Data.Title, r.Data.Owner.Name, r.Data.Stat.View, r.Data.Stat.Danmaku, r.Data.Stat.Like, r.Data.Stat.Reply, r.Data.Stat.Coin, r.Data.Stat.Favorite)
	_, err = f.WriteString(info)
	catch_error(err)
}
