.PHONY: build clean

get-dep:
	go mod download

build-all: build-bff build-cb build-pa build-pt build-pts

build-bff: get-dep
	@echo "-> building module: bff-state-and-aggregate"
	env GOOS=linux go build -ldflags="-s -w" -o bff-state-and-aggregate/bin/agg-read-api bff-state-and-aggregate/agg-read-api/*.go
	env GOOS=linux go build -ldflags="-s -w" -o bff-state-and-aggregate/bin/agg-readall-api bff-state-and-aggregate/agg-readall-api/*.go
	env GOOS=linux go build -ldflags="-s -w" -o bff-state-and-aggregate/bin/state-read-api bff-state-and-aggregate/state-read-api/*.go
	env GOOS=linux go build -ldflags="-s -w" -o bff-state-and-aggregate/bin/state-update-api bff-state-and-aggregate/state-update-api/*.go

build-cb: get-dep
	@echo "-> building module: collect-buzzwords"
	env GOOS=linux go build -ldflags="-s -w" -o collect-buzzwords/bin/collect collect-buzzwords/collect/*.go

build-pa: get-dep
	@echo "-> building module: persist-aggregates"
	env GOOS=linux go build -ldflags="-s -w" -o persist-aggregates/bin/persist persist-aggregates/persist/*.go
	env GOOS=linux go build -ldflags="-s -w" -o persist-aggregates/bin/readall persist-aggregates/readall/*.go
	env GOOS=linux go build -ldflags="-s -w" -o persist-aggregates/bin/read persist-aggregates/read/*.go
	env GOOS=linux go build -ldflags="-s -w" -o persist-aggregates/bin/readall-api persist-aggregates/readall-api/*.go
	env GOOS=linux go build -ldflags="-s -w" -o persist-aggregates/bin/read-api persist-aggregates/read-api/*.go

build-pt: get-dep
	@echo "-> building module: poll-tweet"
	env GOOS=linux go build -ldflags="-s -w" -o poll-tweet/bin/poll poll-tweet/poll/*.go

build-pts: get-dep
	@echo "-> building module: poll-tweet-state"
	env GOOS=linux go build -ldflags="-s -w" -o poll-tweet-state/bin/delete-api poll-tweet-state/delete-api/*.go
	env GOOS=linux go build -ldflags="-s -w" -o poll-tweet-state/bin/read poll-tweet-state/read/*.go
	env GOOS=linux go build -ldflags="-s -w" -o poll-tweet-state/bin/read-api poll-tweet-state/read-api/*.go
	env GOOS=linux go build -ldflags="-s -w" -o poll-tweet-state/bin/update poll-tweet-state/update/*.go
	env GOOS=linux go build -ldflags="-s -w" -o poll-tweet-state/bin/update-api poll-tweet-state/update-api/*.go

clean-pa:
	rm -rf ./persist-aggregates/bin
