# libcoins

gomobile demo for libcoins


## Dependencies

* [Android NDK](https://developer.android.com/ndk)
* [gradle](https://gradle.org)
* [gomobile](https://github.com/golang/go/wiki/Mobile)
* [dep](https://github.com/golang/dep), because of [golang/go#27234](https://github.com/golang/go/issues/27234)

## CLI

    $ make build && ./demo

## Android

Run emulator:

    $ emulator -avd Pixel

Build & deploy:

    $ cd app/android && make bind build install

## iOS

Build & deploy:

    $ cd app/ios && make bind build

## Web

TBD
