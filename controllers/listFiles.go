package controllers

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func ListFiles(client *s3.Client, bucket string) (*s3.ListObjectsV2Output, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: &bucket,
	}
	output, err := client.ListObjectsV2(context.TODO(), input)
	if err != nil {
		log.Printf("Failed to list files: %v", err)
		return nil, fmt.Errorf("failed to retrieve files")
	}
	for _, object := range output.Contents {
		fmt.Println("File: for ", *object.Key)
	}
	return output, nil
}
