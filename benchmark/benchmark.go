package benchmark

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rohankarn35/aws-golang/config"
	"github.com/rohankarn35/aws-golang/controllers"
	"github.com/rohankarn35/aws-golang/models"
	prometheusmetrics "github.com/rohankarn35/aws-golang/prometheus_metrics"
	"github.com/rohankarn35/aws-golang/services/generators"
)

// Run performs a one-time benchmark and returns results
func Run(cfg config.Config) []models.S3SpeedTestResult {
	// Generate files once
	for fileName, size := range cfg.FileSizes {
		fullPath := filepath.Join("static", fileName)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			generators.GenerateFile(fileName, size)
		}
	}

	var results []models.S3SpeedTestResult
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i, region := range cfg.Regions {
		wg.Add(1)
		go func(i int, region string) {
			defer wg.Done()

			client, err := config.GetS3Client(region, cfg.AccessKey, cfg.SecretKey)
			if err != nil {
				log.Printf("Failed to load config for region %s: %v", region, err)
				return
			}

			bucket := cfg.BucketNames[i]
			_, err = client.HeadBucket(context.TODO(), &s3.HeadBucketInput{Bucket: &bucket})
			if err != nil {
				log.Printf("Bucket %s in region %s is inaccessible: %v", bucket, region, err)
				return
			}

			for fileName, size := range cfg.FileSizes {
				s3FileName := filepath.Join("static", fileName)
				for _, sc := range cfg.StorageClass {
					result := benchmarkS3(client, bucket, s3FileName, size, region, sc)
					mu.Lock()
					results = append(results, result)
					// Set initial metrics
					if result.UploadTimeMs > 0 {
						fileSizeMB := float64(size) / (1024 * 1024)
						prometheusmetrics.UploadSpeed.WithLabelValues(region, sc).Set(fileSizeMB / (result.UploadTimeMs / 1000))
						log.Printf("Set initial UploadSpeed for %s/%s: %.2f MB/s", region, sc, fileSizeMB/(result.UploadTimeMs/1000))
					}
					if result.RetrieveTimeMs > 0 {
						fileSizeMB := float64(size) / (1024 * 1024)
						prometheusmetrics.DownloadSpeed.WithLabelValues(region, sc).Set(fileSizeMB / (result.RetrieveTimeMs / 1000))
						log.Printf("Set initial DownloadSpeed for %s/%s: %.2f MB/s", region, sc, fileSizeMB/(result.RetrieveTimeMs/1000))
					}
					if result.DeleteTimeMs > 0 {
						prometheusmetrics.DeleteLatency.WithLabelValues(region).Set(result.DeleteTimeMs / 1000)
						log.Printf("Set initial DeleteLatency for %s: %.2f s", region, result.DeleteTimeMs/1000)
					}
					mu.Unlock()
				}
			}
		}(i, region)
	}

	wg.Wait()
	return results
}

// UpdateMetricsPeriodically simulates ongoing metric updates
func UpdateMetricsPeriodically(cfg config.Config) {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		var wg sync.WaitGroup
		for i, region := range cfg.Regions {
			wg.Add(1)
			go func(i int, region string) {
				defer wg.Done()

				client, err := config.GetS3Client(region, cfg.AccessKey, cfg.SecretKey)
				if err != nil {
					log.Printf("Failed to load config for region %s: %v", region, err)
					return
				}

				bucket := cfg.BucketNames[i]
				_, err = client.HeadBucket(context.TODO(), &s3.HeadBucketInput{Bucket: &bucket})
				if err != nil {
					log.Printf("Bucket %s in region %s is inaccessible: %v", bucket, region, err)
					return
				}

				for fileName, size := range cfg.FileSizes {
					s3FileName := filepath.Join("static", fileName)
					for _, sc := range cfg.StorageClass {
						result := benchmarkS3(client, bucket, s3FileName, size, region, sc)
						if result.UploadTimeMs > 0 {
							fileSizeMB := float64(size) / (1024 * 1024)
							prometheusmetrics.UploadSpeed.WithLabelValues(region, sc).Set(fileSizeMB / (result.UploadTimeMs / 1000))
							log.Printf("Updated UploadSpeed for %s/%s: %.2f MB/s", region, sc, fileSizeMB/(result.UploadTimeMs/1000))
						}
						if result.RetrieveTimeMs > 0 {
							fileSizeMB := float64(size) / (1024 * 1024)
							prometheusmetrics.DownloadSpeed.WithLabelValues(region, sc).Set(fileSizeMB / (result.RetrieveTimeMs / 1000))
							log.Printf("Updated DownloadSpeed for %s/%s: %.2f MB/s", region, sc, fileSizeMB/(result.RetrieveTimeMs/1000))
						}
						if result.DeleteTimeMs > 0 {
							prometheusmetrics.DeleteLatency.WithLabelValues(region).Set(result.DeleteTimeMs / 1000)
							log.Printf("Updated DeleteLatency for %s: %.2f s", region, result.DeleteTimeMs/1000)
						}
					}
				}
			}(i, region)
		}
		wg.Wait()
	}
}

func benchmarkS3(client *s3.Client, bucket, fileName string, size int, region, storageClass string) models.S3SpeedTestResult {
	var uploadTime, retrieveTime, deleteTime time.Duration

	uploadTime = controllers.MeasureExecutionTime(func() {
		err := controllers.UploadFile(client, bucket, fileName)
		if err != nil {
			log.Printf("Upload failed for %s in %s (storage class %s): %v", fileName, region, storageClass, err)
			prometheusmetrics.OperationErrors.WithLabelValues("upload", region).Inc()
		}
	})

	if uploadTime > 0 {
		retrieveTime = controllers.MeasureExecutionTime(func() {
			controllers.RetrieveFile(client, bucket, fileName)
		})
		if retrieveTime == 0 {
			log.Printf("Retrieve skipped or failed for %s in %s (storage class %s)", fileName, region, storageClass)
			prometheusmetrics.OperationErrors.WithLabelValues("retrieve", region).Inc()
		}
	}

	if uploadTime > 0 {
		deleteTime = controllers.MeasureExecutionTime(func() {
			err := controllers.DeleteFile(client, bucket, fileName)
			if err != nil {
				log.Printf("Delete failed for %s in %s (storage class %s): %v", fileName, region, storageClass, err)
				prometheusmetrics.OperationErrors.WithLabelValues("delete", region).Inc()
			}
		})
	}

	fileSizeStr := fmt.Sprintf("%dMB", size/(1024*1024))
	return models.S3SpeedTestResult{
		Region:         region,
		BucketName:     bucket,
		FileSize:       fileSizeStr,
		StorageClass:   storageClass,
		UploadTimeMs:   float64(uploadTime.Milliseconds()),
		RetrieveTimeMs: float64(retrieveTime.Milliseconds()),
		DeleteTimeMs:   float64(deleteTime.Milliseconds()),
	}
}
