# libcoins

gomobile demo for libcoins


## Dependencies

Go:

* Go 1.12, see ([issue #929](https://github.com/gopherjs/gopherjs/issues/929))
* [dep](https://github.com/golang/dep), see [golang/go#27234](https://github.com/golang/go/issues/27234)
* [gomobile](https://github.com/golang/go/wiki/Mobile)
* [gopherjs](https://github.com/gopherjs/gopherjs)

Android:

* [Android NDK](https://developer.android.com/ndk)
* [gradle](https://gradle.org)

## CLI

    $ make build && ./demo

## Android

Run emulator:

    $ emulator -avd Pixel

Build & deploy:

    $ cd app/android && make bind build install

## iOS

Run simulator:

    $ open /Applications/Xcode.app/Contents/Developer/Applications/Simulator.app

Build & deploy:

    $ cd app/ios && make bind build

## Web

TBD
