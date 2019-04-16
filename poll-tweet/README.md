# Serverless by Example

## Solution Overview

## AWS Keys
### AWS Console
Login: https://680260899871.signin.aws.amazon.com/console
User: oliver.koeth@nttdata.com.developer
Password: See LastPass

### AWS APIs
See LastPass => AWS serverless-demo
```
export AWS_ACCESS_KEY_ID=xxx
export AWS_SECRET_ACCESS_KEY=xxx
```

To manage credentials with serverless 
```
sls config credentials --provider aws --key xxx --secret xxx --profile serverless-demo-dev
``` 

### Twitter APIs
See LastPass => Twitter okoeth-demo-consumer / Twitter okoeth-demo-access
```
export TWITTER_ACCESS_TOKEN=xxx
export TWITTER_ACCESS_SECRET=xxx
export TWITTER_CONSUMER_KEY=xxx
export TWITTER_CONSUMER_SECRET=xxx
```

To add credentials to AWS
```
aws secretsmanager create-secret --name serverless-demo/twitter/access-token --region eu-central-1 --secret-string xxx
aws secretsmanager create-secret --name serverless-demo/twitter/access-secret --region eu-central-1 --secret-string xxx
aws secretsmanager create-secret --name serverless-demo/twitter/consumer-key --region eu-central-1 --secret-string xxx
aws secretsmanager create-secret --name serverless-demo/twitter/consumer-secret --region eu-central-1 --secret-string xxx

aws secretsmanager get-secret-value --secret-id serverless-demo/twitter/access-token --region eu-central-1

aws secretsmanager delete-secret --secret-id serverless-demo/twitter/access-token --region eu-central-1
```

For local testing copy `setenv-twitter-template.sh` to `setenv-twitter.sh` and fill in your API keys. Then source the variables:
```
source setenv-twitter.sh
```

## Getting Started
Configure which region / stage / profile to use:
```
source setenv-dev.sh
```

For hello-world:
```
sls create -t aws-go-dep -p poll-tweet
cd poll-tweet
make
sls deploy --region eu-central-1 --stage dev --aws-profile serverless-demo-dev
```

For echo:
```
curl -X POST https://76atmory14.execute-api.eu-central-1.amazonaws.com/dev/echo -d 'Hello, world!'
sls logs -f echo
```

For poll:
```
source setenv-twitter.sh
```

```
aws s3 cp state.json s3://dev-serverless-demo-twitter-state/
```

## References
[Serverless - AWS Documentation](https://serverless.com/framework/docs/providers/aws/)
[Serverless Lambda Go Events](https://serverless.com/blog/framework-example-golang-lambda-support/)
[AWS Lambda Go Events](https://github.com/aws/aws-lambda-go/tree/master/events)
[Twitter Event Source](https://github.com/awslabs/aws-serverless-twitter-event-source)
[IAM Configuration](https://gist.github.com/ServerlessBot/7618156b8671840a539f405dea2704c8)
[Language Analysis](https://github.com/chrisport/go-lang-detector)
[Sentiment Analysis](https://github.com/cdipaolo/sentiment)
