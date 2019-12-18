#!/usr/bin/env bash
export SERVERLESS_STAGE= #dev
export SERVERLESS_USER= #okoeth
export AWS_REGION= #eu-central-1
export AWS_PROFILE=sls-demo-${SERVERLESS_USER}-dev
export AWS_ACCOUNT_ID= #878568643968
export AWS_EVENT_STREAM_NAME_POOL_TWEET=events-poll-tweet-${SERVERLESS_USER}
export AWS_EVENT_STREAM_NAME_COLLECT_BUZZWORDS=events-collect-buzzwords-${SERVERLESS_USER}
export AWS_EVENT_STREAM_POLL_TWEET_ARN="arn:aws:kinesis:"$AWS_REGION":"$AWS_ACCOUNT_ID":stream/"$AWS_EVENT_STREAM_NAME_POOL_TWEET
export AWS_EVENT_STREAM_COLLECT_BUZZWORDS_ARN="arn:aws:kinesis:"$AWS_REGION":"$AWS_ACCOUNT_ID":stream/"$AWS_EVENT_STREAM_NAME_COLLECT_BUZZWORDS
export TWITTER_STATE_FILE=TwitterState.json
export TWITTER_STATE_BUCKET=${SERVERLESS_STAGE}-sls-demo-twitter-state-${SERVERLESS_USER}
export TWITTER_STATE_TABLE=sls-demo-twitter-state
export TWITTER_STATE_TYPE=db
export BUZZWORD_AGGREGATES_TABLE=sls-demo-buzzwords-aggregates
