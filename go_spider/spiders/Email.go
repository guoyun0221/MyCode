package spiders

import "fmt"
import "net/http"
import "io/ioutil"
import "regexp"

func Email() {
	/* To get some emails in a page */

	fmt.Println("Input the link you wanna crawl:")
	var url string //get url
	fmt.Scanln(&url)

	expr := "[\\w]+?@[\\w]+?\\.[\\w]+" // Regular Expression

	res, err := http.Get(url)
	catch_error(err)

	body, err := ioutil.ReadAll(res.Body)
	catch_error(err)
	defer res.Body.Close()
	html := string(body)

	re := regexp.MustCompile(expr)
	list := re.FindAllStringSubmatch(html, -1)

	for _, p := range list {
		fmt.Println(p[0]) //p[0] for all string, p[1] for the stuff in the first (),
	}
}
