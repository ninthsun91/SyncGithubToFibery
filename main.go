package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ninthsun91/SyncGithubToFibery/request"
	"github.com/ninthsun91/SyncGithubToFibery/utils"
)

type LambdaEvent struct {
	Body    string            `json:"body"`
	Headers map[string]string `json:"headers"`
}

func HandleRequest(event *LambdaEvent) (*string, error) {
	signature := event.Headers["x-hub-signature-256"]
	if !utils.VerifySignature(signature, event.Body) {
		fmt.Println("invalid signature")
		return nil, fmt.Errorf("invalid signature")
	}

	err := request.Sync()
	if err != nil {
		return nil, err
	}
	fmt.Println("sync started")

	check := true
	for check {
		resp, err := request.Status()
		if err != nil {
			fmt.Println("error getting status", err)
			check = false
			return nil, err
		}

		if resp.State != request.IN_PROGRESS {
			check = false
			fmt.Println("sync completed")
			return &resp.Message, nil
		}

		fmt.Println("sync in progress")
		time.Sleep(time.Second * 5)
	}

	return nil, fmt.Errorf("unknown error")
}

func main() {
	lambda.Start(HandleRequest)
}
