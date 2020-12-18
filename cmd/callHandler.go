package cmd

import (
	"../pkg"
)

func Call(url string, method string, body string, headers []string) {
	if url != "" && (method == "" || method == "GET") && body == "" && len(headers) == 0 {
		pkg.CallGet(url)
	} else if (url != "" && (method == "" || method == "GET") && body == "" && len(headers) != 0) || (url != "" && method == "GET" && body == "" && len(headers) != 0) {
		pkg.CallGetWithHeader(url, headers)
	}
}
