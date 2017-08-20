default:
	go build -o server
	GOOS=linux GOARCH=arm go build -o server-arm