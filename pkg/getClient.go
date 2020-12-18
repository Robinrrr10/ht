package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

func CallGetWithHeader(url string, headers []string) {
	//fmt.Println("URL:", url)
	req, err := http.NewRequest("GET", url, nil)
	for _, head := range headers {
		headKeyValue := strings.Split(head, ":")
		//fmt.Println("Header:", headKeyValue[0], headKeyValue[1])
		req.Header.Set(headKeyValue[0], headKeyValue[1])
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println("STATUS CODE:", resp.StatusCode)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
