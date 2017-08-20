default: desktop
	GOOS=linux GOARCH=arm GOARM=5 go build -o server-arm

desktop:
	go build -o server