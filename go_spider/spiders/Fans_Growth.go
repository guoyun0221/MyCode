package spiders

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type response struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Ttl     int       `json:"ttl"`
	Data    user_info `json:"data"`
}

type user_info struct {
	Mid       int `json:"mid"`
	Following int `json:"following"`
	Whisper   int `json:"whisper"`
	Black     int `json:"black"`
	Follower  int `json:"follower"`
}

func Fans_Growth() {
	fmt.Println("请输入要查询的up主id")
	var s string
	fmt.Scanln(&s)

	n0 := get_follower(s)
	t0 := time.Now()
	var n1 int
	var t1 time.Time
	fmt.Println("当前粉丝数：", n0, "   当前时间：", t0)

	for i := 0; i < 6; i++ {
		time.Sleep(time.Second * 10)
		n1 = get_follower(s)
		t1 = time.Now()
		fmt.Println("当前粉丝数：", n1, "   当前时间：", t1)
	}

	t := time.Since(t0).Seconds()
	n := get_follower(s) - n0
	v := float64(n) / (t / 60)
	fmt.Println("---------------------------------------")
	fmt.Printf("粉丝增速：%.2f/min\n", v)
}

func get_follower(id string) int {
	url := "https://api.bilibili.com/x/relation/stat?vmid=" + id

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	catch_error(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:69.0) Gecko/20100101 Firefox/69.0")
	resp, err := client.Do(req)
	catch_error(err)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var re response
	err = json.Unmarshal(body, &re)
	catch_error(err)

	fans := re.Data.Follower
	return fans
}
