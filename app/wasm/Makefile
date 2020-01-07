GO111MODULE=off

build:
	GOOS=js GOARCH=wasm go build -o libcoins.wasm

install:
	go get -u github.com/shurcooL/goexec

start:
	goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'

.PHONY: build install start
