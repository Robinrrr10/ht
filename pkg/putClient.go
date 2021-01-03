package pkg

import (
	"net/http"
	"strings"
)

func CallPut(url string, body string) {
	//fmt.Println("URL:", url)
	ioBody := strings.NewReader(body)
	req, err := http.NewRequest("PUT", url, ioBody)
	client := &http.Client{}
	resp, err := client.Do(req)
	responseHandler(resp, err)
}

func CallPutWithHeaders(url string, body string, headers []string) {
	//fmt.Println("URL:", url)
	ioBody := strings.NewReader(body)
	req, err := http.NewRequest("PUT", url, ioBody)
	for _, head := range headers {
		headKeyValue := strings.Split(head, ":")
		//fmt.Println("Header:", headKeyValue[0], headKeyValue[1])
		req.Header.Set(headKeyValue[0], headKeyValue[1])
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	responseHandler(resp, err)
}
