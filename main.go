package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func handler(ctx context.Context, s3Event events.S3Event) error {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	if err != nil {
		fmt.Printf("Failed to initialize session %v ", err)
	}

	svc := s3.New(sess)

	for _, records := range s3Event.Records {
		bucket := records.S3.Bucket.Name
		key := records.S3.Object.Key

		resp, err := svc.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})

		if err != nil {
			fmt.Errorf("Failed to fetch records %v ", err)
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Errorf("Failed to read records %v ", err)
		}

		fmt.Printf("Contents of %s: \n%s\n ", key, string(body))
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
