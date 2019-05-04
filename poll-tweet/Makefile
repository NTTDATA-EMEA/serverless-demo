.PHONY: init-state init-secrets build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/poll poll/*.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: build
	sls deploy --verbose

init: init-state init-secrets

init-state:
	aws s3 ls s3://${TWITTER_STATE_BUCKET}/ >/dev/null 2>/dev/null ; \
	if [ $$? -eq 255 ]; then \
	  aws s3 mb s3://${TWITTER_STATE_BUCKET}/ ; \
	fi
	aws s3 cp TwitterState.json s3://${TWITTER_STATE_BUCKET}/

init-secrets:
	aws secretsmanager get-secret-value --secret-id serverless-demo/twitter/access-token >/dev/null 2>/dev/null ; \
	if [ $$? -eq 255 ]; then \
		aws secretsmanager create-secret --name serverless-demo/twitter/access-token --secret-string ${TWITTER_ACCESS_TOKEN} ; \
	fi
	aws secretsmanager get-secret-value --secret-id serverless-demo/twitter/access-secret >/dev/null 2>/dev/null ; \
	if [ $$? -eq 255 ]; then \
		aws secretsmanager create-secret --name serverless-demo/twitter/access-secret --secret-string ${TWITTER_ACCESS_SECRET} ; \
	fi
	aws secretsmanager get-secret-value --secret-id serverless-demo/twitter/consumer-key >/dev/null 2>/dev/null ; \
	if [ $$? -eq 255 ]; then \
		aws secretsmanager create-secret --name serverless-demo/twitter/consumer-key --secret-string ${TWITTER_CONSUMER_KEY} ; \
	fi
	aws secretsmanager get-secret-value --secret-id serverless-demo/twitter/consumer-secret >/dev/null 2>/dev/null ; \
	if [ $$? -eq 255 ]; then \
		aws secretsmanager create-secret --name serverless-demo/twitter/consumer-secret --secret-string ${TWITTER_CONSUMER_SECRET} ; \
	fi

