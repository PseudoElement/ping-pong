start:
	make build-mac && ./builds/ping-pong-macos

build-mac:
	GOOS=darwin GOARCH=arm64 go build -o ./builds/ping-pong-macos && cp -R assets builds/

build-win:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o ./builds/ping-pong-win.exe && cp -R assets builds/
