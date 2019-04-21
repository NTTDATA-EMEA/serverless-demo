# Poll Tweets

## Function Overview

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

