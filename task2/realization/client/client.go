package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CallVersion(client *http.Client, ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8081/version", nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}

func CallDecode(client *http.Client, ctx context.Context, input string) (string, error) {
	data := map[string]string{"inputString": input}
	jsonData, _ := json.Marshal(data)

	req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:8081/decode", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	return result["outputString"], nil // Вернем декодированную строку
}

func CallHardOp(client *http.Client, ctx context.Context) (bool, int) {
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8081/hard-op", nil)
	if err != nil {
		log.Fatalf("Error creating /hard-op request: %v", err)
		return false, 0
	}

	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Request timed out")
			return false, 500
		}
		log.Fatalf("Error in /hard-op: %v", err)
		return false, 0
	}
	defer resp.Body.Close()

	return true, resp.StatusCode
}
