# Poll Tweet State

Defines two lambda functions to read and update the state used to query Twitter.

## Function Overview 

### Read Function

Reads the state from the DynamoDB. 

The state is defined as a map with twitter's hashtag as key and the id of the latest tweet found with for this hashtag.
```(go)
type State map[string]int64
```

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

Updates the state in the DynamoDB.

## Build and Deploy

### Environment Setup

To prepare proper serverless configuration there is a need to set some environment variables. 
Please use [`setenv-dev-template.sh`](../setenv-dev-template.sh) and fill it with proper values as described [here](../README.md).

### Handler Build

Use Makefile to build executables with help of the following command:
```
$ make build
```
After successful build you will find two files in the `./bin` directory. 

### Handler Deployment

Use Makefile to deploy executables to AWS with help of the following command:
```
$ make deploy
```
After successful deployment you will find two lambda functions in your AWS dashboard together with a relevant DynamoDB table.

### Initial Setup

Before you start using our serverless demo you have to initialize the database with the default state values using Update lambda with e.g. the following payload:

```
{
  "#ai": 0,
  "#cloud": 0,
  "#iot": 0,
  "#startups": 0,
  "#tesla": 0
}
```
The simplest way to initialize the database is to use test functionality of a lambda found in the AWS dashboard.
Just create an new test event with the payload above and run the test. It will write the provided payload in the database.
To check if writing a the state was successful you can run a test for the read lambda. 
It will load the state from the database and show its value in logs.
