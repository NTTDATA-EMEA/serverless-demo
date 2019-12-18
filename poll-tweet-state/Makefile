.PHONY: build clean deploy

build:
	go mod download
	env GOOS=linux go build -ldflags="-s -w" -o bin/read read/*.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/update update/*.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/delete delete/*.go

clean:
	rm -rf ./bin

deploy: build
	sed -e 's/__SERVERLESS_USER__/'${SERVERLESS_USER}'/g' < serverless-template.yml > serverless.yml
	sls deploy --verbose
