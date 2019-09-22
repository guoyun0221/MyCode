package spiders

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func download_pic(pic_url string) {

	fmt.Println(pic_url)
	resp, err := http.Get(pic_url)
	catch_error(err)

	var picture []byte
	if resp != nil { //to avoid "invalid memory address or nil pointer dereference"
		picture, err = ioutil.ReadAll(resp.Body)
		catch_error(err)
		defer resp.Body.Close()
	}

	pic_name := "pics/" + strconv.Itoa(rand.Intn(99999)) + "" + pic_url[len(pic_url)-5:]
	//use a random number and the last 5 characters of link(like x.jpg) to name the file
	if strings.Index(pic_name, ".") == -1 { // is not .jpg/.png and so on
		pic_name = "pics/" + strconv.Itoa(rand.Intn(99999)) + ".jpg"
	}

	f, err := os.Create(pic_name)
	catch_error(err)
	defer f.Close()

	_, err = f.Write(picture)
	catch_error(err)
}
