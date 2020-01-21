# Poll Tweets

Defines one lambda function to poll the tweets from Twitter using API and write them to the defined Kinesis stream. 

## Function Overview

### Poll Function

Queries Twitter, creates events and pushes them to the Kinesis stream.

```
{ 
  "event":"POLL_TWEET_RUN_QUERY",
  "timestamp":"2020-01-20T17:00:04Z",
  "subject": {
     "buzzword":"#cloud",
     "partition_key":"#cloud"
  },
  "object": {
    "tweet_id":1219303590080077824,
    "tweet_text":"Braking on ice: Mastering #data complexity in winter testing with the #IoT.️ Find out how our Bosch IoT Insights… https://t.co/39TZN2SULI"
  }
}
```

## Build and Deploy

### Environment Setup

To prepare proper serverless configuration there is a need to set some environment variables. 
Please use [`setenv-dev-template.sh`](../setenv-dev-template.sh) and fill it with proper values 
and run as described [here](../README.md).

### Twitter APIs

Additionally we have to setup information to be able to connect to Twitter API.
Please use [`setenv-twitter-template.sh`](./setenv-twitter-template.sh) and it with proper values.

```(sh)
export TWITTER_ACCESS_TOKEN=#please provide a valid value
export TWITTER_ACCESS_SECRET=#please provide a valid value
export TWITTER_CONSUMER_KEY=#please provide a valid value
export TWITTER_CONSUMER_SECRET=#please provide a valid value
```

Finally we have to create secrets with help of the following command:
```(sh)
$ make init
```
which in background runs following commands:
```(sh)
aws secretsmanager create-secret --name sls-demo-$SERVERLESS_USER/twitter/access-token --region $AWS_REGION --secret-string $TWITTER_ACCESS_TOKEN
aws secretsmanager create-secret --name sls-demo-$SERVERLESS_USER/twitter/access-secret --region $AWS_REGION --secret-string $TWITTER_ACCESS_SECRET
aws secretsmanager create-secret --name sls-demo-$SERVERLESS_USER/twitter/consumer-key --region $AWS_REGION --secret-string $TWITTER_CONSUMER_KEY
aws secretsmanager create-secret --name sls-demo-$SERVERLESS_USER/twitter/consumer-secret --region $AWS_REGION --secret-string $TWITTER_CONSUMER_SECRET
```

### Handler Build

Use Makefile to build executables with help of the following command:
```(sh)
$ make build
```
After successful build you will find one file `poll` in the `./bin` directory. 

### Handler Deployment

Use Makefile to deploy executables to AWS with help of the following command:
```(sh)
$ make deploy
```
After successful deployment you will find one lambda functions in your AWS dashboard 
together with a relevant Kinesis stream.
