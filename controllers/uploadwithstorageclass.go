package controllers

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// Function to upload a file to S3 with a specific storage class
func UploadFileWithStorageClass(client *s3.Client, bucketName, filePath, storageClass string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to open file, " + err.Error())
	}
	defer file.Close()

	// Get the file name from the file path
	fileName := filepath.Base(filePath)

	// Upload the file with the specified storage class
	uploader := manager.NewUploader(client)
	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:       aws.String(bucketName),
		Key:          aws.String(fileName),
		Body:         file,
		StorageClass: types.StorageClass(storageClass),
	})

	if err != nil {
		log.Fatal("Unable to upload file, " + err.Error())
	}

	fmt.Printf("File uploaded successfully: %s (Storage Class: %s)\n", fileName, storageClass)
}
