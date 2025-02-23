package generators

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// GenerateFile creates a file of the specified size (in bytes) in the static directory
func GenerateFile(fileName string, sizeInBytes int) {
	// Define the static directory and full file path
	staticDir := "static"
	filePath := filepath.Join(staticDir, fileName)

	// Check if the file already exists in the static directory
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		fmt.Printf("File already exists: %s\n", filePath)
		return // Do nothing if the file exists
	}

	// Check if the static directory exists
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		// If static directory doesn't exist, create it
		err = os.Mkdir(staticDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create static directory: %v", err)
		}
		fmt.Printf("Created static directory: %s\n", staticDir)
	}

	// Create the file in the static directory
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	// Create a buffer of random bytes
	buffer := make([]byte, sizeInBytes)
	_, err = rand.Read(buffer)
	if err != nil {
		log.Fatalf("Failed to generate random data: %v", err)
	}

	// Write the buffer to the file
	_, err = file.Write(buffer)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

	fmt.Printf("File generated successfully: %s (Size: %d bytes)\n", filePath, sizeInBytes)
}
