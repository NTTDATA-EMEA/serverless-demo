.PHONY: build clean

build-pa:
	go mod download
	env GOOS=linux go build -ldflags="-s -w" -o persist-aggregates/bin/persist persist-aggregates/persist/*.go
	env GOOS=linux go build -ldflags="-s -w" -o persist-aggregates/bin/readall persist-aggregates/readall/*.go
	env GOOS=linux go build -ldflags="-s -w" -o persist-aggregates/bin/readall-api persist-aggregates/readall-api/*.go
	env GOOS=linux go build -ldflags="-s -w" -o persist-aggregates/bin/read persist-aggregates/read/*.go
	env GOOS=linux go build -ldflags="-s -w" -o persist-aggregates/bin/read-api persist-aggregates/read-api/*.go

clean-pa:
	rm -rf ./persist-aggregates/bin
