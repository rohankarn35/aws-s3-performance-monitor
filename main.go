package main

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"github.com/rohankarn35/aws-golang/benchmark"
)

// var fileSizes = map[string]int{
// 	"1MB":   1 * 1024 * 1024,
// 	"10MB":  10 * 1024 * 1024,
// 	"100MB": 100 * 1024 * 1024,
// 	"1GB":   1024 * 1024 * 1024,
// }

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// generators.GenerateFile("1MB.txt", 1*1024*1024)   // 1MB
	// generators.GenerateFile("10MB.txt", 10*1024*1024) // 10MB
	// generators.GenerateFile("100MB.txt", 100*1024*1024)

	// Get AWS credentials from environment variables
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region1 := os.Getenv("AWS_REGION_1")
	region2 := os.Getenv("AWS_REGION_2")
	region3 := os.Getenv("AWS_REGION_3")
	bucketName1 := os.Getenv("AWS_BUCKET_NAME_1")
	bucketName2 := os.Getenv("AWS_BUCKET_NAME_2")
	bucketName3 := os.Getenv("AWS_BUCKET_NAME_3")

	if accessKey == "" || secretKey == "" || region1 == "" {
		log.Fatalf("Missing AWS credentials or region in environment variables")
	}

	regions := []string{region1, region2, region3}
	bucketNames := []string{bucketName1, bucketName2, bucketName3}
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cfg, err := config.LoadDefaultConfig(context.TODO(),
				config.WithRegion(regions[i]),
				config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
			)
			if err != nil {
				log.Fatalf("Failed to load configuration for region %s: %v", regions[i], err)
			}

			// Create S3 client
			client := s3.NewFromConfig(cfg)
			log.Printf("Created S3 client for region %s and bucket %s", regions[i], bucketNames[i])

			storageClasses := []string{"STANDARD", "STANDARD_IA", "ONEZONE_IA", "INTELLIGENT_TIERING"}
			fileSizes := map[string]string{
				"static/1MB.txt":   "1MB",
				"static/10MB.txt":  "10MB",
				"static/100MB.txt": "100MB",
			}
			benchmark.UploadBenchmark(client, bucketNames[i], storageClasses, fileSizes)
		}(i)
	}

	wg.Wait()

	// cfg, err := config.LoadDefaultConfig(context.TODO(),
	// 	config.WithRegion(region1),
	// 	config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	// )
	// if err != nil {
	// 	log.Fatalf("Failed to load configuration: %v", err)
	// }

	// Create S3 client
	// client := s3.NewFromConfig(cfg)
	// storageClasses := []string{
	// 	"STANDARD",
	// 	"STANDARD_IA",
	// 	"ONEZONE_IA",
	// 	"INTELLIGENT_TIERING",
	// }

	// var result models.S3SpeedTestResult

	// for _, sc := range storageClasses {
	// 	fmt.Printf("\nUploading with Storage Class: %s\n", sc)
	// 	uploadDuration := controllers.MeasureExecutionTime(func() {
	// 		controllers.UploadFileWithStorageClass(client, bucketName1, "file.txt", sc)

	// 	})

	// 	log.Printf("The upload duration for %s is %v", sc, uploadDuration)
	// }

	// Set up test: one test per fileSize, with three regions per test.
	// numFiles := len(fileSizes)
	// results := make([]models.S3SpeedTestResult, numFiles*3)
	// var wg sync.WaitGroup
	// wg.Add(numFiles)

	// i := 0
	// for _, size := range fileSizes {
	// 	go func(i int, size int) {
	// 		defer wg.Done()
	// 		results[i*3] = testUploadAndRetrieveSpeed(accessKey, secretKey, region1, bucketName1, size)
	// 		results[i*3+1] = testUploadAndRetrieveSpeed(accessKey, secretKey, region2, bucketName2, size)
	// 		results[i*3+2] = testUploadAndRetrieveSpeed(accessKey, secretKey, region3, bucketName3, size)
	// 	}(i, size)
	// 	i++
	// }
	// wg.Wait()

	// // Generate PDF with the test results
	// services.GeneratePDF(results)
}

// func testUploadAndRetrieveSpeed(accessKey, secretKey, region, bucketName string, fileSize int) models.S3SpeedTestResult {
// 	// Load AWS configuration with custom credentials
// 	cfg, err := config.LoadDefaultConfig(context.TODO(),
// 		config.WithRegion(region),
// 		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
// 	)
// 	if err != nil {
// 		log.Fatalf("Failed to load configuration: %v", err)
// 	}

// 	// Create S3 client
// 	client := s3.NewFromConfig(cfg)

// 	// Generate file content of size fileSize bytes
// 	fileName := "large_test_file.txt"
// 	fileContent := make([]byte, fileSize)
// 	for i := range fileContent {
// 		fileContent[i] = 'A'
// 	}

// 	// Measure upload time
// 	uploadDuration := controllers.MeasureExecutionTime(func() {
// 		controllers.UploadFile(client, bucketName, fileName, fileContent)
// 	})
// 	fmt.Printf("Upload time for bucket %s in region %s: %v\n", bucketName, region, uploadDuration)

// 	// Measure retrieve time
// 	retrieveDuration := controllers.MeasureExecutionTime(func() {
// 		controllers.RetrieveFile(client, bucketName, fileName)
// 	})
// 	fmt.Printf("Retrieve time for bucket %s in region %s: %v\n", bucketName, region, retrieveDuration)

// 	// Measure delete time
// 	deleteDuration := controllers.MeasureExecutionTime(func() {
// 		err := controllers.DeleteFile(client, bucketName, fileName)
// 		if err != nil {
// 			log.Printf("Failed to delete file: %v", err)
// 		}
// 	})
// 	fmt.Printf("Delete time for bucket %s in region %s: %v\n", bucketName, region, deleteDuration)

// 	var fileSizeStr string
// 	for key, value := range fileSizes {
// 		if value == fileSize {
// 			fileSizeStr = key
// 			break
// 		}
// 	}

// 	return models.S3SpeedTestResult{
// 		Region:         region,
// 		BucketName:     bucketName,
// 		FileSize:       fileSizeStr,
// 		UploadTimeMs:   float64(uploadDuration.Milliseconds()),
// 		RetrieveTimeMs: float64(retrieveDuration.Milliseconds()),
// 		DeleteTimeMs:   float64(deleteDuration.Milliseconds()),
// 	}
// }
