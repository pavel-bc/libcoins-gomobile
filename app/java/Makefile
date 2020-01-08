all: build compile exec

build:
	go build -buildmode=c-shared -o libcoins.so

clean:
	rm -rf ./*.class

compile:
	javac -cp ".:jna-5.5.0.jar" Main.java
	mv ./*.class libcoins/

exec:
	java -cp ".:jna-5.5.0.jar" libcoins.Main

.PHONY: all build clean compile exec
