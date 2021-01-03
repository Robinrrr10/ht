package cmd

func giveRequestArguments(args []string) (url string, method string, body string, headers []string) {
	eachHeaderCount := 0
	for inx, arg := range args {
		if arg == "-u" {
			url = args[inx+1]
		}
		if arg == "-m" {
			method = args[inx+1]
		} else if arg == "GET" || arg == "POST" || arg == "PUT" {
			method = arg
		}
		if arg == "-b" {
			body = args[inx+1]
		}
		if arg == "-h" {
			headers = append(headers, args[inx+1])
			eachHeaderCount++
		}
	}
	if url == "" {
		for inx, arg := range args {
			if arg == "GET" || arg == "POST" || arg == "PUT" {
				url = args[inx+1]
			}
		}
	}
	if len(args) >= 2 && url == "" && args[1] != "-u" && args[1] != "-m" && args[1] != "-b" && args[1] != "-h" {
		url = args[1]
	}
	return
}
