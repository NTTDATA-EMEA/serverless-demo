# Collect Buzzwords

Defines one lambda function which aggregates tweets got from one Kinesis stream and publish the aggregates into another stream.

## Function Overview

### Collect Function

Waits for events published by `poll-tweet` module in Kinesis stream.

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

The handler collects these events and creates a new aggregated event for each buzzwords.

```
{
  "event": "COLLECT_BUZZWORDS_AGGREGATED",
  "timestamp": "0001-01-01T00:00:00Z",
  "subject": {
    "partition_key": "#iot",
    "keyword": "#iot"
  },
  "object": {
    "aggregated_buzzwords": {
      "keyword": "#iot",
      "buzzwords": {
        "#5G": {
          "keyword": "#iot",
          "buzzword": "#5G",
          "count": 1,
          "last_update": "2020-01-21T17:36:27.660281234Z"
        },
        "#AI,": {
          "keyword": "#iot",
          "buzzword": "#AI,",
          "count": 1,
          "last_update": "2020-01-21T17:36:27.660300862Z"
        },
        "#Blockchain": {
          "keyword": "#iot",
          "buzzword": "#Blockchain",
          "count": 1,
          "last_update": "2020-01-21T17:36:27.660301512Z"
        },
        "#IoT,": {
          "keyword": "#iot",
          "buzzword": "#IoT,",
          "count": 1,
          "last_update": "2020-01-21T17:36:27.660281651Z"
        },
        "#privacy": {
          "keyword": "#iot",
          "buzzword": "#privacy",
          "count": 1,
          "last_update": "2020-01-21T17:36:27.660284525Z"
        }
      }
    }
  }
}
```

## Build and Deploy

### Environment Setup

To prepare proper serverless configuration there is a need to set some environment variables. 
Please use [`setenv-dev-template.sh`](../setenv-dev-template.sh) and fill it with proper values 
and run as described [here](../README.md).

### Handler Build

Use Makefile to build executables with help of the following command:
```(sh)
$ make build
```
After successful build you will find one file `collect` in the `./bin` directory. 

### Handler Deployment

Use Makefile to deploy executables to AWS with help of the following command:
```(sh)
$ make deploy
```
After successful deployment you will find one lambda functions in your AWS dashboard 
together with a relevant Kinesis stream.
