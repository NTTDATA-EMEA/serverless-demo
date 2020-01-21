.PHONY: build clean deploy

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

