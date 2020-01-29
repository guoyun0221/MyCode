package spiders

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

type question struct {
	content string
	follow  int
	watch   int
	link    string
}

func Zhihu_question() {
	var start, end int
	fmt.Println("请输入开始爬取的问题id(就是问题网址链接最后面那串数字)，爬虫会从这里开始对id递增依次爬取")
	fmt.Scanf("%d", &start)
	end = start + 1000000

	for i := start; i < end; i++ {
		var q question
		q.link = "https://www.zhihu.com/question/" + strconv.Itoa(i)
		crawl_question(q)
		time.Sleep(time.Second / 20)
	}
}

func crawl_question(q question) {
	expr_question := `class="QuestionHeader-title">([\s\S]+?)<`
	expr_attention := `class="NumberBoard-itemValue"[\s]title="([\d]+?)">[\S]+?</strong>`

	client := &http.Client{}
	req, err := http.NewRequest("GET", q.link, nil)
	catch_error(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:69.0) Gecko/20100101 Firefox/69.0")
	resp, err := client.Do(req)
	catch_error(err)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	html := string(body)

	re_question := regexp.MustCompile(expr_question) //try to find question content
	ques := re_question.FindAllStringSubmatch(html, -1)

	if len(ques) != 0 { // skip blank page，zhihu's question id is not continous
		q.content = ques[0][1]

		re_attention := regexp.MustCompile(expr_attention)
		attention := re_attention.FindAllStringSubmatch(html, -1) //attention[0]:follow, attention[1]:watch
		q.follow, err = strconv.Atoi(attention[0][1])
		catch_error(err)
		q.watch, err = strconv.Atoi(attention[1][1])
		catch_error(err)

		fmt.Printf("问题: %s   关注: %d 浏览: %d   链接: %s\n", q.content, q.follow, q.watch, q.link)
		save_question(q)
	}
}

func save_question(q question) {
	f, err := os.OpenFile("Zhihu_question.txt", os.O_APPEND, 0666) //save data
	catch_error(err)
	defer f.Close()
	info := fmt.Sprintf("问题: %s   关注: %d 浏览: %d   链接: %s\n", q.content, q.follow, q.watch, q.link)
	_, er := f.WriteString(info)
	catch_error(er)
}
