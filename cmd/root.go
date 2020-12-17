package cmd

import (
	"fmt"
	"os"
)

func init() {
	//fmt.Println("currently in root")
}

func Execute() {

	cmdParams := os.Args
	if len(cmdParams) <= 1 {
		rootHelp()
	} else {
		url, method, body, headers := giveRequestArguments(cmdParams)
		fmt.Println("url:", url)
		fmt.Println("method:", method)
		fmt.Println("body:", body)
		fmt.Println("headers:", headers)
	}
}

func rootHelp() {
	fmt.Println("Supporting commands:")
	fmt.Println("-m : --method : GET : POST")
	fmt.Println("-u : --url")
	fmt.Println("-h : --header")
	fmt.Println("-b : --body")
}
