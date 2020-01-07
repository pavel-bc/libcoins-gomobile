build:
	go build -buildmode=c-shared -o libcoins.a main.go
	gcc -o demo example.c libcoins.a

.PHONY: build
