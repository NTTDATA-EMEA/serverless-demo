.PHONY: build clean deploy

build:
	go mod download
	env GOOS=linux go build -ldflags="-s -w" -o bin/agg-readall-api agg-readall-api/*.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/agg-read-api agg-read-api/*.go

clean:
	rm -rf ./bin

deploy:
	sed -e 's/__SERVERLESS_USER__/'${SERVERLESS_USER}'/g' < serverless-template.yml > serverless.yml
	sls deploy --verbose
