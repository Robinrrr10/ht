package pkg

import (
	"net/http"
	"strings"
)

func CallPost(url string, body string) {
	//fmt.Println("URL:", url)
	ioBody := strings.NewReader(body)
	req, err := http.NewRequest("POST", url, ioBody)
	client := &http.Client{}
	resp, err := client.Do(req)
	responseHandler(resp, err)
}

func CallPostWithHeaders(url string, body string, headers []string) {
	//fmt.Println("URL:", url)
	ioBody := strings.NewReader(body)
	req, err := http.NewRequest("POST", url, ioBody)
	for _, head := range headers {
		headKeyValue := strings.Split(head, ":")
		//fmt.Println("Header:", headKeyValue[0], headKeyValue[1])
		req.Header.Set(headKeyValue[0], headKeyValue[1])
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	responseHandler(resp, err)
}
