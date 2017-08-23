default: desktop
	GOOS=linux GOARCH=arm GOARM=5 go build -v -o server-arm

desktop:
	go build -v -o server