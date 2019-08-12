.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/read read/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/update update/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/delete delete/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: build
	sed -e 's/__SERVERLESS_USER__/'${SERVERLESS_USER}'/g' < serverless-template.yml > serverless.yml
	sls deploy --verbose
