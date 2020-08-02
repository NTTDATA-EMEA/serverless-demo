.PHONY: init build clean deploy

build:
	go mod download
	env GOOS=linux go build -ldflags="-s -w" -o bin/poll poll/*.go

clean:
	rm -rf ./bin

deploy:
	sed -e 's/__SERVERLESS_USER__/'${SERVERLESS_USER}'/g' < serverless-template.yml > serverless.yml
	sls deploy --verbose

secrets:
	aws secretsmanager get-secret-value --secret-id sls-demo-${SERVERLESS_USER}/twitter/access-token >/dev/null 2>/dev/null ; \
	if [ $$? -eq 255 ]; then \
		aws secretsmanager create-secret --name sls-demo-${SERVERLESS_USER}/twitter/access-token --secret-string ${TWITTER_ACCESS_TOKEN} ; \
	fi
	aws secretsmanager get-secret-value --secret-id sls-demo-${SERVERLESS_USER}/twitter/access-secret >/dev/null 2>/dev/null ; \
	if [ $$? -eq 255 ]; then \
		aws secretsmanager create-secret --name sls-demo-${SERVERLESS_USER}/twitter/access-secret --secret-string ${TWITTER_ACCESS_SECRET} ; \
	fi
	aws secretsmanager get-secret-value --secret-id sls-demo-${SERVERLESS_USER}/twitter/consumer-key >/dev/null 2>/dev/null ; \
	if [ $$? -eq 255 ]; then \
		aws secretsmanager create-secret --name sls-demo-${SERVERLESS_USER}/twitter/consumer-key --secret-string ${TWITTER_CONSUMER_KEY} ; \
	fi
	aws secretsmanager get-secret-value --secret-id sls-demo-${SERVERLESS_USER}/twitter/consumer-secret >/dev/null 2>/dev/null ; \
	if [ $$? -eq 255 ]; then \
		aws secretsmanager create-secret --name sls-demo-${SERVERLESS_USER}/twitter/consumer-secret --secret-string ${TWITTER_CONSUMER_SECRET} ; \
	fi

