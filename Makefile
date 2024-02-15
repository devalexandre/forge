.PHONY: build clean deploy

build:
	rm -rf ./bin
	rm -rf ./vendor
	go get ./...
	go mod vendor
	env GOOS=linux go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o bin/forge ./main.go
	if [ ! -d ~/.local/bin ]; then mkdir -p ~/.local/bin; fi
	cp ./bin/forge ~/.local/bin/forge

clean:
	rm -rf ./bin ./vendor

