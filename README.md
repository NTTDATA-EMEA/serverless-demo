# Serverless Demo

## Solution Overview

## Identity and Access Management

## AWS Keys
The AWS keys are provided in an AWS profile configuration which is provides
```
sls config credentials --provider aws --key xxx --secret xxx --profile serverless-demo-dev
``` 

## IAM Profile
The AWS keys are tied towards the `serverless-demo` user who is backed by a `serverless-demo`
policy. The details of the policy have been included [here](./serverless-demo-policy).

## Appendix A: Cheat Sheets

## Appendix B: References
The following resources proved to be useful during the creation of this demo

### Serverless
* [Serverless AWS Documentation](https://serverless.com/framework/docs/providers/aws/)
* [Serverless Lambda Go Events](https://serverless.com/blog/framework-example-golang-lambda-support/)
* [Serverless Examples](https://github.com/serverless/examples)
* [Serverless IAM Configuration](https://gist.github.com/ServerlessBot/7618156b8671840a539f405dea2704c8)

### AWS
* [AWS SDK for Go](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/welcome.html)
* [AWS Lambda Go Events](https://github.com/aws/aws-lambda-go/tree/master/events)
* [AWS Kinesis Introduction](https://docs.aws.amazon.com/streams/latest/dev/key-concepts.html)
* [AWS Kinesis CLI](https://docs.aws.amazon.com/streams/latest/dev/fundamental-stream.html)
* [AWS S3 CLI]()

### Other
* [Twitter Event Source](https://github.com/awslabs/aws-serverless-twitter-event-source)
* [Language Analysis](https://github.com/chrisport/go-lang-detector)
* [Sentiment Analysis](https://github.com/cdipaolo/sentiment)

