#!/usr/bin/env bash
export TWITTER_ACCESS_TOKEN=#please provide a valid value
export TWITTER_ACCESS_SECRET=#please provide a valid value
export TWITTER_CONSUMER_KEY=#please provide a valid value
export TWITTER_CONSUMER_SECRET=#please provide a valid value
export TWITTER_STATE_FILE=TwitterState.json
export TWITTER_STATE_BUCKET=${SERVERLESS_STAGE}-sls-demo-twitter-state-${SERVERLESS_USER}
export TWITTER_QUERY_TIMEOUT_SEC=10

aws secretsmanager create-secret --name sls-demo-$SERVERLESS_USER/twitter/access-token --region $AWS_REGION --secret-string $TWITTER_ACCESS_TOKEN
aws secretsmanager create-secret --name sls-demo-$SERVERLESS_USER/twitter/access-secret --region $AWS_REGION --secret-string $TWITTER_ACCESS_SECRET
aws secretsmanager create-secret --name sls-demo-$SERVERLESS_USER/twitter/consumer-key --region $AWS_REGION --secret-string $TWITTER_CONSUMER_KEY
aws secretsmanager create-secret --name sls-demo-$SERVERLESS_USER/twitter/consumer-secret --region $AWS_REGION --secret-string $TWITTER_CONSUMER_SECRET
