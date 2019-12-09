.PHONY: init-state init-secrets build clean deploy

build:
	go mod download
	env GOOS=linux go build -ldflags="-s -w" -o bin/poll poll/*.go

clean:
	rm -rf ./bin

deploy: build
	sed -e 's/__SERVERLESS_USER__/'${SERVERLESS_USER}'/g' < serverless-template.yml > serverless.yml
	sls deploy --verbose

init: init-state init-secrets

init-state:
	aws s3 ls s3://${TWITTER_STATE_BUCKET}/ >/dev/null 2>/dev/null ; \
	if [ $$? -eq 255 ]; then \
	  aws s3 mb s3://${TWITTER_STATE_BUCKET}/ ; \
	  aws s3api put-bucket-tagging --bucket ${TWITTER_STATE_BUCKET} --tagging 'TagSet=[{Key=project,Value=sls-demo-'${SERVERLESS_USER}'},{Key=owner,Value='${SERVERLESS_USER}'}]' ; \
	fi
	aws s3 cp TwitterState.json s3://${TWITTER_STATE_BUCKET}/

init-secrets:
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

