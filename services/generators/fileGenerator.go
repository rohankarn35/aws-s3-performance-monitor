package generators

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
)

// GenerateFile creates a file of the specified size (in bytes)
func GenerateFile(fileName string, sizeInBytes int) {
	// Ensure the static directory exists
	staticDir := "static"
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		err = os.Mkdir(staticDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create static directory: %v", err)
		}
	}

	// Create the file in the static directory
	filePath := fmt.Sprintf("%s/%s", staticDir, fileName)
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
