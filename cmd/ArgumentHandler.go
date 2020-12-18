package cmd

func giveRequestArguments(args []string) (url string, method string, body string, headers []string) {
	eachHeaderCount := 0
	for inx, arg := range args {
		if arg == "-u" {
			url = args[inx+1]
		}
		if arg == "-m" {
			method = args[inx+1]
		}
		if arg == "-b" {
			body = args[inx+1]
		}
		if arg == "-h" {
			//headers[eachHeaderCount] = args[inx+1]
			headers = append(headers, args[inx+1])
			eachHeaderCount++
		}
	}
	return
}
