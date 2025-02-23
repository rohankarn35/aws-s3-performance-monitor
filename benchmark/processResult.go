package benchmark

import (
	"fmt"

	"github.com/rohankarn35/aws-golang/models"
	"github.com/rohankarn35/aws-golang/services/generators"
	results_service "github.com/rohankarn35/aws-golang/services/results"
)

func ProcessResults(results []models.S3SpeedTestResult) {
	csvData := [][]string{{"Region", "Bucket", "File Size", "Storage Class", "Upload (ms)", "Retrieve (ms)", "Delete (ms)"}}
	for _, r := range results {
		csvData = append(csvData, []string{
			r.Region, r.BucketName, r.FileSize, r.StorageClass,
			fmt.Sprintf("%.2f", r.UploadTimeMs), fmt.Sprintf("%.2f", r.RetrieveTimeMs), fmt.Sprintf("%.2f", r.DeleteTimeMs),
		})
	}
	results_service.SaveResultsToCSV("s3_benchmark_results.csv", csvData)
	generators.GeneratePDF(results)
}
