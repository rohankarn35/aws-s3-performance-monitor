# AWS S3 Performance Benchmarking Tool

![Go](https://img.shields.io/badge/Go-1.23-blue?logo=go) ![AWS S3](https://img.shields.io/badge/AWS-S3-orange?logo=amazonaws) ![Prometheus](https://img.shields.io/badge/Prometheus-2.51-red?logo=prometheus) ![Grafana](https://img.shields.io/badge/Grafana-10.2-yellow?logo=grafana)

A powerful Go-based tool to benchmark AWS S3 upload, download, and delete performance across multiple regions and storage classes, exposing real-time metrics via Prometheus and visualizing them with Grafana.

---

## Overview

This project measures the performance of AWS S3 operations (upload, download, delete) for various file sizes (1MB, 10MB, 100MB) across multiple regions and storage classes (STANDARD, STANDARD_IA, ONEZONE_IA). It generates detailed reports in CSV and PDF formats, exposes metrics through a Prometheus endpoint, and provides a live Grafana dashboard for visualization—ideal for understanding S3 performance in real-world scenarios.

### Key Features

- **Multi-Region Benchmarking**: Tests S3 performance across configurable regions (e.g., `us-east-1`, `us-west-2`, `eu-west-1`).
- **Storage Class Analysis**: Evaluates different S3 storage classes for latency and throughput.
- **Real-Time Metrics**: Exposes upload/download speeds, delete latency, and error counts via Prometheus.
- **Grafana Visualization**: Displays dynamic graphs for performance monitoring.
- **Report Generation**: Saves results as CSV and PDF for offline analysis.
- **Efficient Design**: Built with Go for speed and concurrency, using goroutines for parallel benchmarking.

---

## Prerequisites

- **Go**: Version 1.23 or higher
- **AWS Account**: With S3 buckets and valid credentials
- **Prometheus**: Version 2.51.0 (or compatible)
- **Grafana**: Version 10.2.4 (or compatible)
- **Git**: For cloning the repository

---

## Installation

### Clone the Repository

```bash
git clone https://github.com/rohankarn35/aws-s3-benchmark.git
cd aws-s3-benchmark
```

### Install Go Dependencies

```bash
go mod tidy
```

### Set Up `.env`

Create a `.env` file in the project root with your AWS credentials and S3 bucket details:

```
AWS_ACCESS_KEY_ID=your_access_key
AWS_SECRET_ACCESS_KEY=your_secret_key
AWS_REGION_1=us-east-1
AWS_REGION_2=us-west-2
AWS_REGION_3=eu-west-1
AWS_BUCKET_NAME_1=your-bucket-1
AWS_BUCKET_NAME_2=your-bucket-2
AWS_BUCKET_NAME_3=your-bucket-3
```

> **Note**: Ensure buckets exist in the specified regions and your AWS user has permissions (e.g., `s3:PutObject`, `s3:GetObject`, `s3:DeleteObject`).

### Download Prometheus

- Get it from [Prometheus Downloads](https://prometheus.io/download/).
- Extract:
  ```bash
  tar xvfz prometheus-2.51.0.linux-amd64.tar.gz
  cd prometheus-2.51.0.linux-amd64
  ```
- Create `prometheus.yml`:
  ```yaml
  global:
          scrape_interval: 15s
  scrape_configs:
          - job_name: 'aws-s3-benchmark'
                  static_configs:
                          - targets: ['localhost:8080']
  ```

### Download Grafana

- Get it from [Grafana Downloads](https://grafana.com/grafana/download).
- Extract:
  ```bash
  tar -zxvf grafana-10.2.4.linux-amd64.tar.gz
  cd grafana-10.2.4
  ```

---

## Usage

Run the application, Prometheus, and Grafana in separate terminals for the full demo experience.

### Step 1: Run the Application

```bash
go run main.go
```

- Benchmarks S3 performance once, generates `s3_benchmark_results.csv` and `s3_performance_report.pdf`.
- Starts a Prometheus metrics server at `http://localhost:8080/metrics`, updating metrics every 15 seconds.

### Step 2: Run Prometheus

```bash
./prometheus --config.file=prometheus.yml
```

- Scrapes metrics from `http://localhost:8080/metrics` every 15 seconds.
- Access the Prometheus UI at `http://localhost:9090`.

### Step 3: Run Grafana

```bash
./bin/grafana-server
```

- Open `http://localhost:3000`, log in with `admin/admin` (change password on first login).
- Add Prometheus as a data source:
  - Configuration > Data Sources > Add > Prometheus.
  - URL: `http://localhost:9090`.
  - Save & Test.
- Create a dashboard:
  - “+” > Dashboard > Add panel.
  - Queries: `s3_upload_speed_mb_per_s`, `s3_download_speed_mb_per_s`, `s3_delete_latency_seconds`, `s3_operation_errors_total`.
  - Save as “S3 Performance Dashboard”.

### Demo Workflow

1. Show the app running and generating logs (e.g., metric updates).
2. Open `http://localhost:8080/metrics` to display raw metrics.
3. Switch to `http://localhost:9090/targets` to confirm Prometheus scraping.
4. Present the Grafana dashboard at `http://localhost:3000` for live graphs.

---

## Makefile

This project includes a `Makefile` to streamline common tasks such as building, running, and cleaning up Docker containers.

### Available Commands

- **Build the Docker image**:

  ```bash
  make build
  ```

- **Run the Docker container**:

  ```bash
  make run
  ```

- **Stop and remove the Docker container**:

  ```bash
  make stop
  ```

- **Clean up Docker images and containers**:

  ```bash
  make clean
  ```

- **Rebuild and run the Docker container**:

  ```bash
  make rerun
  ```

- **Check the logs of the Docker container**:
  ```bash
  make logs
  ```

These commands help manage the Docker environment efficiently, ensuring a smooth development and deployment process.

## Metrics Exposed

| Metric Name                  | Type    | Labels                    | Description                         |
| ---------------------------- | ------- | ------------------------- | ----------------------------------- |
| `s3_upload_speed_mb_per_s`   | Gauge   | `region`, `storage_class` | Upload speed in MB/s                |
| `s3_download_speed_mb_per_s` | Gauge   | `region`, `storage_class` | Download speed in MB/s              |
| `s3_delete_latency_seconds`  | Gauge   | `region`                  | Delete operation latency in seconds |
| `s3_operation_errors_total`  | Counter | `operation`, `region`     | Total errors by operation           |

---

## Sample Output

### Logs

```
2025/02/23 12:34:56 Initial UploadSpeed for us-east-1/STANDARD: 10.23 MB/s
2025/02/23 12:35:11 Updated UploadSpeed for us-east-1/STANDARD: 10.25 MB/s
2025/02/23 12:35:11 Metrics scraped at 2025-02-23T12:35:11Z
```

### Grafana Dashboard

![Sample Grafana Dashboard](https://raw.githubusercontent.com/rohankarn35/rohankarn35/main/Screenshot%20from%202025-02-23%2013-52-18.png)

---

## Why This Project Stands Out

- **Real-World Relevance**: Measures S3 performance critical for cloud applications.
- **Monitoring Integration**: Combines Prometheus and Grafana for professional-grade observability.
- **Extensibility**: Easily adaptable for additional regions, storage classes, or metrics.
- **Clean Code**: Modular design with Go best practices (goroutines, error handling).

---

## Contributing

This is a demo project, but feel free to fork and enhance it! Suggestions:

- Add more S3 operations (e.g., list objects).
- Implement real-time S3 polling instead of simulated updates.
- Support additional cloud providers.

---

## License

MIT License - feel free to use and modify this code for educational purposes.

---
