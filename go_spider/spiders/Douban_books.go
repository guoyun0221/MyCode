package spiders

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

func Douban_books() {
	expr := `title="([\s\S]+?)">?[\s\S]+? <p class="pl">([\s\S]+?)</p>[\s\S]+?<span class="rating_nums">([\s\S]+?)</span>`
	const MaxPage = 10
	for i := 0; i < MaxPage*25; i += 25 {
		url := "https://book.douban.com/top250?start=" + strconv.Itoa(i)
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		catch_error(err)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:69.0) Gecko/20100101 Firefox/69.0")
		resp, err := client.Do(req)
		catch_error(err)
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		html := string(body)

		re := regexp.MustCompile(expr)
		list := re.FindAllStringSubmatch(html, -1)

		for j, _ := range list {
			fmt.Println(i+j, "书名:", list[j][1], "  信息:", list[j][2], "  评分:", list[j][3])
		}
	}
}
