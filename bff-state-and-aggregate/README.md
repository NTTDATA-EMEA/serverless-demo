# BFF State and Aggregate

Defines four lambda functions as BFF (backend for frontend) for a UI client.

## Function Overview 

### API Read Aggregate Function

Reads an aggregate using `persist-aggregates` module based on the provided buzzword.

Aggregate example for `#tesla`:

```
{
  "keyword": "#tesla",
  "buzzwords": {
    "#ElectricCar": {
      "keyword": "#tesla",
      "buzzword": "#ElectricCar",
      "count": 1,
      "last_update": "2020-01-21T18:21:16.716759899Z"
    },
    "#German": {
      "keyword": "#tesla",
      "buzzword": "#German",
      "count": 4,
      "last_update": "2020-01-21T18:21:16.716760511Z"
    },
    "#Tesla…": {
      "keyword": "#tesla",
      "buzzword": "#Tesla…",
      "count": 1,
      "last_update": "2020-01-21T18:21:16.716760228Z"
    },
    "#china": {
      "keyword": "#tesla",
      "buzzword": "#china",
      "count": 4,
      "last_update": "2020-01-21T18:21:16.7167592Z"
    }
  }
}
``` 

### API ReadAll Aggregates Function

Reads all aggregates using `persist-aggregates` module.
It is a test function an should be used with caution due to potentially huge payload to transfer.

### API Read State Function

Reads the state using `poll-tweet-state` module. 

State example:
```(json)
{
    "#ai": 1219303578273288194,
    "#cloud": 1219303590080077824,
    "#iot": 1219303590080077824,
    "#startups": 1219243463201513473,
    "#tesla": 1219097834739257344
}
```

### Update Function

Updates the state using `poll-tweet-state` module.

## Build and Deploy

### Environment Setup

To prepare proper serverless configuration there is a need to set some environment variables. 
Please use [`setenv-dev-template.sh`](../setenv-dev-template.sh) and fill it with proper values 
and run as described [here](../README.md).

### Handler Build

Use Makefile to build executables with help of the following command:
```
$ make build
```
After successful build you will find three files in the `./bin` directory. 

### Handler Deployment

Use Makefile to deploy executables to AWS with help of the following command:
```
$ make deploy
```
After successful deployment you will find four lambda functions in your AWS dashboard.
