package main

import (
	"fmt"
	"log"
	"task2/realization/client"
	"time"
)

func main() {
	totalTimeout := 15 * time.Second

	apiClient := client.NewAPIClient("http://localhost:8082", totalTimeout)

	version, err := apiClient.CallVersion()
	if err != nil {
		log.Fatalf("Error in /version: %v", err)
	}
	fmt.Printf("Version: %s\n", version)

	decoded, err := apiClient.CallDecode("SGVsbG8gd29ybGQ=")
	if err != nil {
		log.Fatalf("Error in /decode: %v", err)
	}
	fmt.Printf("Decoded string: %s\n", decoded)

	success, status := apiClient.CallHardOp()
	if !success {
		log.Printf("Hard operation failed with status code: %d\n", status)
	} else {
		fmt.Printf("Hard operation succeeded with status code: %d\n", status)
	}
}
