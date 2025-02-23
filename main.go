package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rohankarn35/aws-golang/benchmark"
	"github.com/rohankarn35/aws-golang/config"
	"github.com/rohankarn35/aws-golang/metrics"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Run benchmarks once
	results := benchmark.Run(cfg)

	// Start metrics server and periodic updates
	go metrics.StartServer()
	go benchmark.UpdateMetricsPeriodically(cfg)

	// Process initial benchmark results
	benchmark.ProcessResults(results)

	log.Println("Benchmarking complete. Results saved to s3_benchmark_results.csv and s3_performance_report.pdf")

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Received shutdown signal, exiting...")
}
