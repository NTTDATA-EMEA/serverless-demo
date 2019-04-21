.PHONY: build test clean

build:
	dep ensure -v
	go build ./..

test:
	go test ./..

clean:
	rm -rf ./vendor Gopkg.lock
