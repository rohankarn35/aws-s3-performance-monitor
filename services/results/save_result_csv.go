package results_service

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func SaveResultsToCSV(fileName string, data [][]string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed to create CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range data {
		err := writer.Write(record)
		if err != nil {
			log.Fatalf("Failed to write to CSV file: %v", err)
		}
	}

	fmt.Printf("Benchmark results saved to %s\n", fileName)
}
