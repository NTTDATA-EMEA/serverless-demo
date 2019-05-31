# Serverless Demo

This is a demo for serverless computing. The demo implements an event streaming / sourcing
pattern using serverless components from AWS. The lock-in to AWS is controlled by the
[Serverless.com](https://serverless.com) framework and also the APIs are wrapped by a
[pkg/services](https://github.com/okoeth/serverless-demo/tree/master/commons/pkg/services) package,
which provides also a local file-system based implementation for testing purposes (and the
demonstration of independence from AWS).

## Solution Overview

![Diagram](https://drive.google.com/uc?authuser=0&id=1PxxyRGoLRru2y8I9RnnU0nVsQHXMdMlQ&export=download)

The functionality implemented by the solution is rather simple: tweets with certain search
keyword-based specification (S3) are queried by a function (Lambda) and published on an event stream (Kinesis).
The event stream (Kinesis) drives another function (Lambda) which processes the events and
generates statistics for buzzwords related to the original keywords. The statistics and the
processing state of the event stream are persited in a document store (DynamoDB). The Twitter
API keys are managed as secrets (SecretManager) and diagnoctic information is made available (CloudWatch).

### Module: poll-tweets

Polls the tweets from Twitter. For more information see: [./poll-tweet/README.md](./poll-tweet/README.md)

### Module: collect buzzwords

Builds a statistic for buzzwords. For more information see: [./collect-buzzwords/README.md](./collect-buzzwords/README.md)

## Prerequisites

The development environment requires a number of packages which are provided in convenient Dockerfile.
For more information see: [./collect-buzzwords/README.md](./collect-buzzwords/README.md)

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

### Using a containerised development environment

To build the containerised development environment run:

```(sh)
docker build -t okoeth/serverless-demo-dev .
```

Run a developer shell with:

```(sh)
docker run -t -i -v "$PWD":/work -v "$HOME":/root --rm okoeth/serverless-demo-dev
```

### Setting the environment

Set environment with:

```(sh)
cd /work
source setenv-dev.sh
```

And make sure the AWS CLI has the proper profile enabled. For the initial set-up an AWS user
dedicated to the serverless-demo should be created. (See below.)

## Identity and Access Management

We will be setting up a dedicated user for the sewrverless demo. We assume that in your current
configuration you are running with an AWS profile that has sufficient privileges to execute
the below IAM commands.

### IAM Policy

The AWS keys used in this demo shall be tied towards the `sls-demo-${SERVERLESS_USER}` user who is backed by a
`sls-demo-${SERVERLESS_USER}-policy` policy. The details of the policy have been included
[here](./serverless-demo-policy). The policy can be created with the following command:

```(sh)
aws iam create-policy --policy-name sls-demo-${SERVERLESS_USER}-policy --policy-document file://serverless-demo-policy.json
```

Then the user can be created with the following command:

```(sh)
aws iam create-user \
  --user-name sls-demo-${SERVERLESS_USER}
  
aws iam attach-user-policy \
  --user-name sls-demo-${SERVERLESS_USER} \
  --policy-arn arn:aws:iam::${AWS_ACCOUNT_ID}:policy/sls-demo-${SERVERLESS_USER}-policy
```

Finally, create the access keys using:

```(sh)
aws iam create-access-key --user-name sls-demo-${SERVERLESS_USER}
```

The AWS keys are provided in an AWS profile configuration which is provides

```(sh)
sls config credentials --provider aws --key xxx --secret xxx --profile sls-demo-${SERVERLESS_USER}-dev
```

Unfortunately, `sls config` cannot set the AWS region, so we run the AWS CLI and provide èu-central-1` as region.

```(sh)
aws configure
```

## Build and Deplyoment

### Commons Package

Make sure you have sources the basic environment with `source setenv-dev.sh` then `cd ./commons`. Initiialise
common resources with:

```(sh)
make init
```

Then resolve dependencies with

```(sh)
make build
```

Note that since we do not create a binary or shared library or anything for commons, there is no actual
`go build` running.

### Poll Package

Then `cd ../poll-tweet`and create resources with:

```(sh)
make init
```

## Appendix A: Cheat Sheets

### Upgrade AWS CLI on MacOS

```(sh)
sudo -H pip install --upgrade awscli --ignore-installed six
```

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
