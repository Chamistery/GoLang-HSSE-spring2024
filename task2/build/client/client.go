package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	client "task2/realization/client"
	"time"
)

func main() {
	curr_client := &http.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	version, err := client.CallVersion(curr_client, ctx)
	if err != nil {
		log.Fatalf("Error in /version: %v", err)
	}
	fmt.Printf("%s\n", version)

	decoded, err := client.CallDecode(curr_client, ctx, "SGVsbG8gd29ybGQ=")
	if err != nil {
		log.Fatalf("Error in /decode: %v", err)
	}
	fmt.Printf("%s\n", decoded)

	success, status := client.CallHardOp(curr_client, ctx)
	fmt.Printf("%v, %d\n", success, status)
}
