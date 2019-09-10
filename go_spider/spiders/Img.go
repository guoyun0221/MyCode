package spiders

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
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

func download_pic(pic string) {
	fmt.Println(pic)
	resp, err := http.Get(pic)
	catch_error(err)

	picture, err := ioutil.ReadAll(resp.Body)
	catch_error(err)
	defer resp.Body.Close()

	f, err := os.Create("pics/" + strconv.Itoa(rand.Intn(99999)) + "" + pic[len(pic)-5:])
	//use a random number and the last 5 characters of link(like x.jpg) to name the file
	catch_error(err)
	defer f.Close()

	_, er := f.Write(picture)
	catch_error(er)
}
