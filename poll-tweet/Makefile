.PHONY: init build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/poll poll/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/echo echo/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

init-state:
	aws s3 cp state.json s3://dev-serverless-demo-twitter-state/ --region ${AWS_REGION}
