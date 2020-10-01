package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, er := http.Get("http://api.theysaidso.com/qod")
	if er != nil {
		fmt.Println(er)
		return
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(contents))
}
