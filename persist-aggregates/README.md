# Persist Aggregates

Defines three lambda functions to persist and read aggregates for buzzwords.

## Function Overview 

### Read Function

Reads an aggregate from the DynamoDB based on the buzzword.

Aggregate example:

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

### ReadAll Function

Reads all aggregates from the DynamoDB.
It is a test function an should be used with caution due to potentially huge payload to transfer.

### Persist Function

Persists an aggregate event from the Kinesis stream in DynamoDB.

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
After successful deployment you will find three lambda functions `read`, `readall` and `persist` in your AWS dashboard together with a relevant DynamoDB table.
