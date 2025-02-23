package generators

import (
	"fmt"

	"log"

	"github.com/jung-kurt/gofpdf"
	"github.com/rohankarn35/aws-golang/models"
)

// Generate PDF report
func GeneratePDF(results []models.S3SpeedTestResult) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "B", 12)
	pdf.AddPage()
	pdf.Cell(190, 10, "S3 Performance Test Results")
	pdf.Ln(10)

	// Table Headers
	pdf.SetFont("Arial", "B", 10)
	headers := []string{"Region", "Bucket", "File Size", "Upload (ms)", "Retrieve (ms)", "Delete (ms)"}
	colWidths := []float64{40, 40, 25, 30, 30, 30}
	for i, h := range headers {
		pdf.CellFormat(colWidths[i], 10, h, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Table Content
	pdf.SetFont("Arial", "", 10)
	for _, result := range results {
		pdf.CellFormat(40, 10, result.Region, "1", 0, "C", false, 0, "")
		pdf.CellFormat(40, 10, result.BucketName, "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 10, result.FileSize, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, fmt.Sprintf("%.2f", result.UploadTimeMs), "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, fmt.Sprintf("%.2f", result.RetrieveTimeMs), "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, fmt.Sprintf("%.2f", result.DeleteTimeMs), "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}

	err := pdf.OutputFileAndClose("s3_performance_report.pdf")
	if err != nil {
		log.Fatalf("Error saving PDF: %v", err)
	}
}
