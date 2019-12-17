.PHONY: build clean deploy init init-stream

build:
	go mod download
	env GOOS=linux go build -ldflags="-s -w" -o bin/persist persist/*.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/readall readall/*.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/read read/*.go

clean:
	rm -rf ./bin

deploy:
	sed -e 's/__SERVERLESS_USER__/'${SERVERLESS_USER}'/g' < serverless-template.yml > serverless.yml
	sls deploy --verbose

init: init-stream

init-stream:
	aws kinesis describe-stream --stream-name ${AWS_EVENT_STREAM_NAME_COLLECT_BUZZWORDS} >/dev/null 2>/dev/null ; \
	if [ $$? -eq 255 ]; then \
		aws kinesis create-stream --stream-name ${AWS_EVENT_STREAM_NAME_COLLECT_BUZZWORDS} --shard-count 1 ; \
		aws kinesis add-tags-to-stream --stream-name ${AWS_EVENT_STREAM_NAME_COLLECT_BUZZWORDS} --tags project=sls-demo-${SERVERLESS_USER},owner=${SERVERLESS_USER}
	fi
