BASE=github.com/pavel-bc/libcoins
PACKAGE=$(BASE)/libcoins

bind:
	CGO_ENABLED=1 gomobile bind -v -target=android -o app/libcoins.aar $(PACKAGE)

build:
	cd app && gradle build

install:
	adb install app/build/outputs/apk/debug/app-debug.apk

.PHONY: bind build install
