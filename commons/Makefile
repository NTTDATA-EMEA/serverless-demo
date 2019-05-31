.PHONY: init-events build test clean

init: init-events

init-events:
	aws kinesis create-stream --stream-name ${AWS_EVENT_STREAM_NAME} --shard-count 1
	aws kinesis add-tags-to-stream --stream-name ${AWS_EVENT_STREAM_NAME} --tags project=sls-demo-${SERVERLESS_USER},owner=${SERVERLESS_USER}

build:
	dep ensure -v

clean: clean-events
	rm -rf ./vendor Gopkg.lock

clean-events:
	aws kinesis delete-stream --stream-name ${AWS_EVENT_STREAM_NAME}
