package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func responseHandler(resp *http.Response, err error) {
	//fmt.Println("-----response-------")
	if err != nil {
		fmt.Println("ERROR:::")
		fmt.Println(err)
	} else {
		fmt.Println("STATUS CODE:::", resp.StatusCode)
		printHeader(resp)
		printBody(resp)
	}
}

func printHeader(resp *http.Response) {
	fmt.Println("HEADERS:::")
	headers := resp.Header
	for key, eachHead := range headers {
		for _, val := range eachHead {
			fmt.Println(key, ":", val)
		}
	}
}

func printBody(resp *http.Response) {
	fmt.Println("BODY:::")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
