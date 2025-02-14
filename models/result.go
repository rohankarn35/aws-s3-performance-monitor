package models

type S3SpeedTestResult struct {
	Region         string
	BucketName     string
	FileSize       string
	UploadTimeMs   float64
	RetrieveTimeMs float64
	DeleteTimeMs   float64
}
