BASE=github.com/pavel-bc/libcoins
PACKAGE=$(BASE)/libcoins

bind:
	gomobile bind -v -target=ios -o Libcoins.framework $(PACKAGE)

build:
	xcodebuild -project bind.xcodeproj -scheme bind -configuration 'Debug' -sdk iphonesimulator

install:
	xcrun simctl install booted DerivedData/bind/Build/Products/Debug-iphonesimulator/bind.app

.PHONY: bind build install
