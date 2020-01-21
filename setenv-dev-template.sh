#!/usr/bin/env bash
export SERVERLESS_STAGE=#please provide a valid value
export SERVERLESS_USER=#please provide a valid value
export AWS_REGION=#please provide a valid value
export AWS_PROFILE=sls-demo-${SERVERLESS_USER}-dev
export AWS_ACCOUNT_ID=#please provide a valid value
export AWS_EVENT_STREAM_NAME_POLL_TWEET=events-poll-tweet-${SERVERLESS_USER}-${SERVERLESS_STAGE}
export AWS_EVENT_STREAM_NAME_COLLECT_BUZZWORDS=events-collect-buzzwords-${SERVERLESS_USER}-${SERVERLESS_STAGE}
export AWS_EVENT_STREAM_ARN_POLL_TWEET="arn:aws:kinesis:"$AWS_REGION":"$AWS_ACCOUNT_ID":stream/"$AWS_EVENT_STREAM_NAME_POOL_TWEET
export AWS_EVENT_STREAM_ARN_COLLECT_BUZZWORDS="arn:aws:kinesis:"$AWS_REGION":"$AWS_ACCOUNT_ID":stream/"$AWS_EVENT_STREAM_NAME_COLLECT_BUZZWORDS
export TWITTER_STATE_FILE=TwitterState.json
export TWITTER_STATE_BUCKET=${SERVERLESS_STAGE}-sls-demo-twitter-state-${SERVERLESS_USER}
export TWITTER_STATE_TABLE=sls-demo-twitter-state
export TWITTER_STATE_TYPE=db
export BUZZWORD_AGGREGATES_TABLE=sls-demo-buzzwords-aggregates
