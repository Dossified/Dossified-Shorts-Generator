CHROMIUM_PATH ?= /usr/bin/chromium

gowitness-server:
	./gowitness server --debug --chrome-path ${CHROMIUM_PATH} --fullpage -X 1080 -Y 1920

run:
	go run main.go

build:
	go build .

format:
	golines -w "."

