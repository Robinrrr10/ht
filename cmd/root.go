package cmd

import (
	"fmt"
	"os"

	"../pkg"
)

func init() {
	//fmt.Println("currently in root")
}

func Execute() {
	//fmt.Println("currently in Execute")
	cmdParams := os.Args
	if len(cmdParams) <= 1 {
		rootHelp()
	} else {
		switch cmdParams[1] {
		case "-m":
			{
				fmt.Println("method")
				switch cmdParams[2] {
				case "GET":
					{
						//fmt.Println("method: GET")
						//fmt.Println("Get method")
						if cmdParams[3] == "" {
							fmt.Println("Please give url after GET")
						}
						pkg.CallGet(cmdParams[3])
					}
				case "POST":
					{
						fmt.Println("method: POST")
					}
				default:
					fmt.Println("invalid method. Use GET or POST")
				}
			}

		case "GET":
			{
				//fmt.Println("Get method")
				if cmdParams[2] == "" {
					fmt.Println("Please give url after GET")
				}
				pkg.CallGet(cmdParams[2])
			}
		default:
			{
				fmt.Println("directly call get with url")
				pkg.CallGet(cmdParams[1])
			}
		}
	}
	//rootHelp()
}

func rootHelp() {
	fmt.Println("Supporting commands:")
	fmt.Println("-m : --method : GET : POST")
	fmt.Println("-u : --url")
	fmt.Println("-h : --header")
	fmt.Println("-b : --body")
}
