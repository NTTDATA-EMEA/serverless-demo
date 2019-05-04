.PHONY: build clean deploy init init-stream

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/collect collect/*.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

init: init-stream

init-stream:
	aws kinesis describe-stream --stream-name ${AWS_EVENT_STREAM_NAME} >/dev/null 2>/dev/null ; \
	if [ $$? -eq 255 ]; then \
		aws kinesis create-stream --stream-name ${AWS_EVENT_STREAM_NAME} --shard-count 1 ; \
	fi
