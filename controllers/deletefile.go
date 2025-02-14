package controllers

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func DeleteFile(client *s3.Client, bucket, key string) error {
	input := &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}
	_, err := client.DeleteObject(context.TODO(), input)
	if err != nil {
		log.Printf("Failed to delete file: %v", err)
		return err
	}
	log.Print("file deleted sucessfully")

	return nil
}
