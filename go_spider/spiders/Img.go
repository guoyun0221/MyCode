package spiders

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"time"
)

func Img() {
	/* To get some pics  */

	rand.Seed(time.Now().Unix())
	var url string
	fmt.Println("Input the link you wanna crawl:")
	fmt.Scanln(&url)
	expr := "<img[\\s\\S]+?src=\"(http[\\s\\S]+?)\""

	res, err := http.Get(url)
	catch_error(err)

	body, err := ioutil.ReadAll(res.Body)
	catch_error(err)
	defer res.Body.Close()
	html := string(body)

	re := regexp.MustCompile(expr)
	list := re.FindAllStringSubmatch(html, -1)

	for _, pic := range list {
		download_pic(pic[1])
	}
}
