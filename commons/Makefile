.PHONY: init-events build test clean

init-events:
	aws kinesis create-stream --stream-name ${AWS_EVENT_STREAM_NAME} --shard-count 1

build:
	dep ensure -v
	go build ./..

test:
	go test ./..

clean:
	rm -rf ./vendor Gopkg.lock

clean-events:
	aws kinesis delete-stream --stream-name ${AWS_EVENT_STREAM_NAME}
