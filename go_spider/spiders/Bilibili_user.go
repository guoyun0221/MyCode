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

type result struct {
	Code    int    `json:"code"`    //it's not ' that near enter key, it's ` below esc key
	Message string `json:"message"` //only capitalized can be transform to json
	Ttl     int    `json:"ttl"`
	Data    user   `json:"data"`
}

type user struct {
	Id    int    `json:"mid"`
	Name  string `json:"name"`
	Sex   string `json:"sex"`
	Level int    `json:"level"`
	Face  string `json:"face"`
}

func Bilibili_user() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000; i++ {
		mid := rand.Intn(450000000) + 1 //450000000 is the approximate range of id
		url := "https://api.bilibili.com/x/space/acc/info?mid=" + strconv.Itoa(mid) + "&jsonp=jsonp"
		get_info(url)
		time.Sleep(time.Second / 10) //don't visit too ofen
	}
}

func get_info(url string) {

	client := &http.Client{} //These below is to avoid 403
	req, err := http.NewRequest("GET", url, nil)
	catch_error(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:69.0) Gecko/20100101 Firefox/69.0")
	resp, err := client.Do(req)
	catch_error(err)

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var res result
	err = json.Unmarshal(body, &res) //trans json to struct
	catch_error(err)

	person := res.Data
	fmt.Printf("id:%d  name:%s  sex:%s  lv:%d  picture:%s\n", person.Id, person.Name, person.Sex, person.Level, person.Face)

	f, err := os.OpenFile("Bilibili_users.txt", os.O_APPEND, 0666) //save data
	catch_error(err)
	defer f.Close()
	info := fmt.Sprintf("id:%d  name:%s  sex:%s  lv:%d  picture:%s\n", person.Id, person.Name, person.Sex, person.Level, person.Face)
	_, er := f.WriteString(info)
	catch_error(er)
}
