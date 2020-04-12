package spiders

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

func Stupid_repeater() {
	fmt.Println("请输入原句：")
	//Scanln can't read sentences with spaces
	reader := bufio.NewReader(os.Stdin)
	data, _, err := reader.ReadLine()
	catch_error(err)
	s := string(data)

	s_en := doTranslate(s)
	fmt.Println("我来复述一下：")
	fmt.Println(doTranslate(s_en))
}

func doTranslate(src string) string {
	expr := `"tgt":"([\s\S]+?)"}`
	URL := "http://fanyi.youdao.com/translate?smartresult=dict&smartresult=rule"
	var dst string

	resp, err := http.PostForm(URL, url.Values{
		"type":        {"AUTO"},
		"i":           {src},
		"from":        {"AUTO"},
		"to":          {"AUTO"},
		"smartresult": {"dict"},
		"client":      {"fanyideskweb"},
		"salt":        {"15627225943846"},
		"sign":        {"06f4f1e1f1f89dcda3bafa1e1ef233df"},
		"ts":          {"1562722594384"},
		"bv":          {"3a019e7d0dda4bcd253903675f2209a5"},
		"doctype":     {"json"},
		"version":     {"2.1"},
		"keyfrom":     {"fanyi.web"},
		"action":      {"FY_BY_CLICKBUTTION"}})
	catch_error(err)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	json := string(body) //I don't know how to decode json array, so I use regexp
	re := regexp.MustCompile(expr)
	str_list := re.FindAllStringSubmatch(json, -1)
	for i := 0; i < len(str_list); i++ {
		dst = dst + str_list[i][1]
	}

	return dst
}
