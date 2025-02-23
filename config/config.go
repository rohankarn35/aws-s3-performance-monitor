package config

import (
	"log"
	"os"
)

type Config struct {
	AccessKey    string
	SecretKey    string
	Regions      []string
	BucketNames  []string
	FileSizes    map[string]int
	StorageClass []string
}

func LoadConfig() (Config, error) {
	// err := godotenv.Load()
	// if err != nil {
	// 	return Config{}, err
	// }

	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	regions := []string{os.Getenv("AWS_REGION_1"), os.Getenv("AWS_REGION_2"), os.Getenv("AWS_REGION_3")}
	bucketNames := []string{os.Getenv("AWS_BUCKET_NAME_1"), os.Getenv("AWS_BUCKET_NAME_2"), os.Getenv("AWS_BUCKET_NAME_3")}

	if accessKey == "" || secretKey == "" || regions[0] == "" {
		log.Fatalf("Missing AWS credentials or region in environment variables")
	}

	// Define file sizes to test
	fileSizes := map[string]int{
		"1MB.txt": 1 * 1024 * 1024,
	}

	// Storage classes to test
	storageClasses := []string{"STANDARD", "STANDARD_IA"}

	return Config{
		AccessKey:    accessKey,
		SecretKey:    secretKey,
		Regions:      regions,
		BucketNames:  bucketNames,
		FileSizes:    fileSizes,
		StorageClass: storageClasses,
	}, nil
}
