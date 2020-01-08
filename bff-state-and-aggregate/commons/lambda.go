package commons

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	//	POLL_TWEET_STATE_MODULE = "poll-tweet-state"
	PERSIST_AGGREGATES_MODULE = "persist-aggregates"
)

func Invoke(mn string, fn string, payload interface{}) (*lambda.InvokeOutput, error) {
	var p []byte
	var err error
	log.Infof("Invoking '%s' with payload: %t", fn, payload)
	if payload != nil {
		switch v := payload.(type) {
		//		case string:
		//			p = []byte(v)
		default:
			p, err = json.Marshal(v)
			if err != nil {
				return nil, err
			}
		}
	}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := lambda.New(sess, &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	user := os.Getenv("SERVERLESS_USER")
	stg := os.Getenv("SERVERLESS_STAGE")
	name := mn + "-" + user + "-" + stg + "-" + fn + "-" + user
	result, err := client.Invoke(&lambda.InvokeInput{
		FunctionName:   aws.String(name),
		InvocationType: aws.String(lambda.InvocationTypeRequestResponse),
		Payload:        p,
	})
	return result, nil
}
