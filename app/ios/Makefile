BASE=github.com/pavel-bc/libcoins
PACKAGE=$(BASE)/libcoins

bind:
	gomobile bind -v -target=ios -o Libcoins.framework $(PACKAGE)

build:
	xcodebuild -project bind.xcodeproj -scheme bind -configuration 'Debug' -sdk iphonesimulator

.PHONY: bind build install
