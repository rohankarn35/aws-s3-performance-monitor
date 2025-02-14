package controllers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func UploadFile(client *s3.Client, bucket, fileName string, storageClass string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error opening file %s", fileName)
	}
	defer file.Close()

	input := &s3.PutObjectInput{
		Bucket:       aws.String(bucket),
		Key:          aws.String(fileName),
		Body:         file,
		StorageClass: types.StorageClass(storageClass),
	}
	_, err = client.PutObject(context.TODO(), input)
	if err != nil {
		log.Printf("Failed to upload file: %v", err)
		return err
	}

	return nil

}
