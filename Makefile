default: desktop
	GOOS=linux GOARCH=arm go build -o server-arm

desktop:
	go build -o server