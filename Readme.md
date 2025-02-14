# AWS Golang Project

This project demonstrates the use of AWS S3 services with Golang. It includes functionalities for uploading, listing, retrieving, and deleting files in S3 buckets. Additionally, it measures the execution time of these operations and generates a PDF report of the performance test results.

## Features

- **File Operations**: Upload, list, retrieve, and delete files in S3 buckets.
- **Execution Time Measurement**: Measure the time taken for each operation.
- **PDF Report Generation**: Generate a PDF report of the performance test results.
- **Storage Class Support**: Upload files with different S3 storage classes.

## Project Structure

- **controllers**: Contains the main logic for interacting with S3.
  - `listFiles.go`: Lists files in an S3 bucket.
  - `deletefile.go`: Deletes a file from an S3 bucket.
  - `measureExecutionTime.go`: Measures the execution time of a function.
  - `retriveFile.go`: Retrieves a file from an S3 bucket.
  - `uploadFile.go`: Uploads a file to an S3 bucket.
  - `uploadwithstorageclass.go`: Uploads a file with a specific storage class.
- **services/generators**: Contains utilities for generating files and PDF reports.
  - `pdfresult_service.go`: Generates a PDF report of the performance test results.
  - `fileGenerator.go`: Generates files of specified sizes.
- **models**: Contains data models.
  - `result.go`: Defines the structure for storing S3 speed test results.
- **static**: Directory for storing generated files.
- **benchmark**: Placeholder for benchmarking scripts.
- **main.go**: Entry point of the application.

## Setup

1. Clone the repository.
2. Create a `.env` file with your AWS credentials and bucket information.
3. Run the application using `go run main.go`.

## Dependencies

- AWS SDK for Go v2
- gofpdf for PDF generation
- godotenv for loading environment variables

## License

This project is licensed under the MIT License.
