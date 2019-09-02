package main

import "fmt"
import "net/http"
import "io/ioutil"
import "os"
import "strconv"

func main() {
	for i := 0; i < 10; i++ {
		var url string = "http://placekitten.com/g/" + strconv.Itoa(320+i*50) + "/" + strconv.Itoa(180+i*50)
		var img string = "../pics/cat_" + strconv.Itoa(320+i*50) + "_" + strconv.Itoa(180+i*50) + ".jpg"

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}

		pic, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		resp.Body.Close()

		f, err := os.Create(img)
		if err != nil {
			fmt.Println(err)
		}

		_, er := f.Write(pic) // it will cause "no new variables on left side of :="
		if er != nil {        // if I still use err
			fmt.Println(err)
		}

		f.Close()
	}
}
