# Serverless Demo

## Solution Overview

## Prerequisites

### Basic Installations

* Golang
* Dep

### Installation of Serverless Framework

```(sh)
sudo npm install serverless -g
```

### Installation of AWS CLI

```(sh)
tbd.
```

## Identity and Access Management

### AWS Keys

The AWS keys are provided in an AWS profile configuration which is provides

```(sh)
sls config credentials --provider aws --key xxx --secret xxx --profile serverless-demo-dev
```

### IAM Profile

The AWS keys are tied towards the `serverless-demo` user who is backed by a `serverless-demo`
policy. The details of the policy have been included [here](./serverless-demo-policy).

## Appendix A: Cheat Sheets

### Read Kinesis via CLI

```(sh)
aws kinesis create-stream --stream-name $AWS_EVENT_STREAM_NAME --shard-count 1

aws kinesis put-record --stream-name $AWS_EVENT_STREAM_NAME --partition-key 123 --data testdata

aws kinesis describe-stream --stream-name $AWS_EVENT_STREAM_NAME --region $AWS_REGION

SHARD_ITERATOR=$(aws kinesis get-shard-iterator --shard-id shardId-000000000000 --shard-iterator-type TRIM_HORIZON --stream-name $AWS_EVENT_STREAM_NAME --query 'ShardIterator')
aws kinesis get-records --shard-iterator $SHARD_ITERATOR
aws kinesis delete-stream --stream-name $AWS_EVENT_STREAM_NAME
```

## Appendix B: References

The following resources proved to be useful during the creation of this demo

### Serverless

* [Serverless AWS Documentation](https://serverless.com/framework/docs/providers/aws/)
* [Serverless Lambda Go Events](https://serverless.com/blog/framework-example-golang-lambda-support/)
* [Serverless Examples](https://github.com/serverless/examples)
* [Serverless IAM Configuration](https://gist.github.com/ServerlessBot/7618156b8671840a539f405dea2704c8)
* [Serverless Localstack](https://github.com/localstack/serverless-localstack)

### AWS

* [AWS SDK for Go](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/welcome.html)
* [AWS Lambda Go Events](https://github.com/aws/aws-lambda-go/tree/master/events)
* [AWS Kinesis Introduction](https://docs.aws.amazon.com/streams/latest/dev/key-concepts.html)
* [AWS Kinesis CLI](https://docs.aws.amazon.com/streams/latest/dev/fundamental-stream.html)
* [AWS S3 CLI](...)

### Other

* [Twitter Event Source](https://github.com/awslabs/aws-serverless-twitter-event-source)
* [Language Analysis](https://github.com/chrisport/go-lang-detector)
* [Sentiment Analysis](https://github.com/cdipaolo/sentiment)
