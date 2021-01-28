package pkg

import (
	"net/http"
	"strings"
)

func CallGet(url string) {
	//fmt.Println("in call get client----")
	resp, err := http.Get(url)
	responseHandler(resp, err)
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
	responseHandler(resp, err)
}
