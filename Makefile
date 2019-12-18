.PHONY: build clean

build-pa:
	go mod download
	env GOOS=linux go build -ldflags="-s -w" -o persist-aggregates/bin/persist persist-aggregates/persist/*.go
	env GOOS=linux go build -ldflags="-s -w" -o persist-aggregates/bin/readall persist-aggregates/readall/*.go
	env GOOS=linux go build -ldflags="-s -w" -o persist-aggregates/bin/read persist-aggregates/read/*.go

clean-pa:
	rm -rf ./persist-aggregates/bin
