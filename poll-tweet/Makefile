.PHONY: init-state build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/poll poll/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

init-state:
	aws s3 ls s3://${TWITTER_STATE_BUCKET}/ --region ${AWS_REGION} >/dev/null 2>/dev/null ; \
	if [ $$? -eq 255 ]; then \
	  aws s3 mb s3://${TWITTER_STATE_BUCKET}/ --region ${AWS_REGION} ; \
	fi
	aws s3 cp TwitterState.json s3://${TWITTER_STATE_BUCKET}/ --region ${AWS_REGION}
