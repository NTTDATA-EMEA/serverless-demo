#!/usr/bin/env bash
export SERVERLESS_STAGE= #dev
export SERVERLESS_USER= #okoeth
export AWS_REGION= #eu-central-1
export AWS_PROFILE=sls-demo-${SERVERLESS_USER}-dev
export AWS_ACCOUNT_ID= #878568643968
export AWS_EVENT_STREAM_NAME=events-${SERVERLESS_USER}
export AWS_EVENT_STREAM_ARN="arn:aws:kinesis:"$AWS_REGION":"$AWS_ACCOUNT_ID":stream/"$AWS_EVENT_STREAM_NAME
