.PHONY: build clean deploy

build:
	go mod download
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/agg-readall-api agg-readall-api/*.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/agg-read-api agg-read-api/*.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/state-read-update-api state-read-update-api/*.go

clean:
	rm -rf ./bin

deploy:
	sed -e 's/__SERVERLESS_USER__/'${SERVERLESS_USER}'/g' < serverless-template.yml > serverless.yml
	sls deploy --verbose
