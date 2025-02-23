package prometheusmetrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	DownloadSpeed = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "s3_download_speed_mb_per_s",
			Help: "Download speed from S3 in MB/s",
		},
		[]string{"region", "storage_class"},
	)
	UploadSpeed = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "s3_upload_speed_mb_per_s",
			Help: "Upload speed to S3 in MB/s",
		},
		[]string{"region", "storage_class"},
	)
	DeleteLatency = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "s3_delete_latency_seconds",
			Help: "Latency of delete operations in seconds",
		},
		[]string{"region"},
	)
	OperationErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "s3_operation_errors_total",
			Help: "Total number of errors encountered during S3 operations",
		},
		[]string{"operation", "region"},
	)
)

func init() {
	prometheus.MustRegister(DownloadSpeed)
	prometheus.MustRegister(UploadSpeed)
	prometheus.MustRegister(DeleteLatency)
	prometheus.MustRegister(OperationErrors)
}
