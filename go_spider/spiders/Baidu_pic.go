package spiders

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func Baidu_pic() {
	rand.Seed(time.Now().Unix())
	expr := `objURL":"(http[\s\S]+?)"`
	var url, key string
	var num, count int
	fmt.Println("输入想要爬取的图片关键字和数量，以空格分隔:")
	fmt.Scanln(&key, &num)

	for pn := 0; count < num; pn += 60 {
		url = "http://image.baidu.com/search/flip?tn=baiduimage&ie=utf-8&word=" + key + "&pn=" + strconv.Itoa(pn) + "&gsm=&ct=&ic=0&lm=-1&width=0&height=0"
		//in the orginal link, change 'index' into 'slip', and turn to next page, then you'll see this kind of link.
		//to trans dynamic loading into page turning. one page contains 20 pics, but it seems like baidu
		//put 3 pages into one source, because pn=0 will get 60 pics
		res, err := http.Get(url)
		catch_error(err)
		body, err := ioutil.ReadAll(res.Body)
		catch_error(err)
		defer res.Body.Close()
		html := string(body)
		re := regexp.MustCompile(expr)
		list := re.FindAllStringSubmatch(html, -1)

		for _, pic := range list {
			count++
			download_pic(pic[1])
			if count == num {
				break
			}
		}
	}
}
