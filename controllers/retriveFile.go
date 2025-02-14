package controllers

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func RetrieveFile(client *s3.Client, bucket, key string) {
	input := &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}
	output, err := client.GetObject(context.TODO(), input)
	if err != nil {
		log.Printf("Failed to retrieve file: %v", err)
		return
	}
	defer output.Body.Close()
	fmt.Println("File retrieved successfully")
}
