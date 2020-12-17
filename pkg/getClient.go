package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CallGet(url string) {
	//fmt.Println("in call get client")
	//fmt.Println("URL:", url)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
