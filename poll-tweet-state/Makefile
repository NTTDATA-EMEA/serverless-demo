.PHONY: build clean deploy

build:
	go mod download
	env GOOS=linux go build -ldflags="-s -w" -o bin/read read/*.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/read-api read-api/*.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/update-api update-api/*.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/delete-api delete-api/*.go

clean:
	rm -rf ./bin

deploy:
	sed -e 's/__SERVERLESS_USER__/'${SERVERLESS_USER}'/g' < serverless-template.yml > serverless.yml
	sls deploy --verbose
