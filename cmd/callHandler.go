package cmd

import (
	"../pkg"
)

func Call(url string, method string, body string, headers []string) {
	if url != "" && (method == "" || method == "GET") && body == "" && len(headers) == 0 {
		pkg.CallGet(url)
	} else if (url != "" && (method == "" || method == "GET") && body == "" && len(headers) != 0) || (url != "" && method == "GET" && body == "" && len(headers) > 0) {
		pkg.CallGetWithHeader(url, headers)
	} else if url != "" && (method == "" || method == "POST") && len(headers) == 0 {
		pkg.CallPost(url, body)
	} else if url != "" && (method == "" || method == "POST") && len(headers) > 0 {
		pkg.CallPostWithHeaders(url, body, headers)
	} else if url != "" && (method == "" || method == "PUT") && len(headers) == 0 {
		pkg.CallPut(url, body)
	} else if url != "" && (method == "" || method == "PUT") && len(headers) > 0 {
		pkg.CallPutWithHeaders(url, body, headers)
	} else if url != "" && (method == "" || method == "DELETE") && len(headers) == 0 {
		pkg.CallDelete(url, body)
	} else if url != "" && (method == "" || method == "DELETE") && len(headers) > 0 {
		pkg.CallDeleteWithHeaders(url, body, headers)
	}
}
