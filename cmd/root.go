package cmd

import "fmt"

func init() {
	//fmt.Println("currently in root")
}

func Execute() {
	//fmt.Println("currently in Execute")
	rooHelp()
}

func rooHelp() {
	fmt.Println("Supporting commands:")
	fmt.Println("-m : --method : GET : POST")
	fmt.Println("-u : --url")
	fmt.Println("-h : --header")
	fmt.Println("-b : --body")
}
