package benchmark

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rohankarn35/aws-golang/controllers"
	results_service "github.com/rohankarn35/aws-golang/services/results"
)

func UploadBenchmark(client *s3.Client, bucketName string, storageClasses []string, fileSizes map[string]string) {
	var wg sync.WaitGroup
	results := [][]string{
		{"FileName", "StorageClass", "FileSize", "Duration", "Speed (Mb/s)", "Status"},
	}

	for fileName, size := range fileSizes {
		for _, storageClass := range storageClasses {
			wg.Add(1)
			go func(fileName, storageClass, size string) {
				defer wg.Done()

				// Start the timer
				startTime := time.Now()

				// Upload the file
				err := controllers.UploadFile(client, bucketName, fileName, storageClass)
				duration := time.Since(startTime).Seconds()
				status := "Success"
				if err != nil {
					status = "Failed"
				}
				// Calculate speed (MB/s)
				fileInfo, err := os.Stat(fileName)
				if err != nil {
					fmt.Printf("Failed to get file info: %s\n", err)

				}
				fileSize := fileInfo.Size()
				fileSizeMb := float64(fileSize * 8 / (1024 * 1024))
				speed := fileSizeMb / (duration + 1e-9)

				// Record the result

				results = append(results, []string{
					fileName, storageClass, size, fmt.Sprintf("%.2f", duration), fmt.Sprintf("%.2f", speed), status,
				})

				fmt.Printf("Upload completed: %s (%s) - Duration: %.2fs, FileSize: %.2fMb Speed: %.2f Mb/s, Status: %s\n",
					fileName, storageClass, duration, fileSizeMb, speed, status)
			}(fileName, storageClass, size)
		}
	}

	// Wait for all uploads to complete
	wg.Wait()

	// Save results to CSV
	csv_filename := fmt.Sprintf("upload_file_test_%s.csv", bucketName)
	results_service.SaveResultsToCSV(csv_filename, results)

}
