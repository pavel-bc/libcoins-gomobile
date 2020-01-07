BASE=github.com/pavel-bc/libcoins
GO111MODULE=off

build:
	go build -o ./demo $(BASE)/app/cli

ensure:
	dep ensure -v

.PHONY: build ensure
