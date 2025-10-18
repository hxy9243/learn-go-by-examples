package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const HeaderKey = "X-Client-Id"

const (
	Success = "SUCCESS"
	Limited = "LIMITED"
	Error   = "ERROR"
	Total   = "Total"
)

func sendRequest(client *http.Client, clientId, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(HeaderKey, clientId)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}

func randomString(n int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rnd.Intn(len(charset))]
	}
	return string(b)

}

// sendRequestsOneClient sends requests every second for determined amount of seconds
func sendRequestsOneClient(clientId, url string, numRequests int, tickerTime, totalTime time.Duration) map[string]int {
	results := map[string]int{
		Success: 0,
		Limited: 0,
		Error:   0,
		Total:   0,
	}

	resultChan := make(chan string)

	startTime := time.Now()
	ticker := time.NewTicker(tickerTime)
	N := 0

	client := &http.Client{}

	for {
		if time.Since(startTime) > totalTime {
			break
		}
		N++
		fmt.Printf("Sending parallel requests for %s: time %d, %d requests in parallel\n", clientId, N, numRequests)

		wg := &sync.WaitGroup{}
		groupStartTime := time.Now()

		for i := 0; i < numRequests; i++ {
			wg.Go(func() {
				resp, err := sendRequest(client, clientId, url)
				if err != nil {
					fmt.Printf("Error sending request to %s: %v\n", url, err)
				}

				switch resp.StatusCode {
				case http.StatusOK:
					resultChan <- Success
				case http.StatusTooManyRequests:
					resultChan <- Limited
				case http.StatusInternalServerError:
					resultChan <- Error
				default:
					resultChan <- Error
				}

			})
		}

		for i := 0; i < numRequests; i++ {
			result := <-resultChan
			results[result]++
			results[Total]++
		}
		wg.Wait()
		fmt.Printf("Client %s group total time: %s\n", clientId, time.Since(groupStartTime))

		<-ticker.C
	}
	return results
}

func main() {
	// parse cli flags
	url := flag.String("url", "http://localhost:8080", "The URL of the service to test")

	flag.Parse()

	fmt.Printf("Testing service at %s\n", *url)

	wg := &sync.WaitGroup{}

	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 100

	for i := 0; i < 10; i++ {
		wg.Go(func() {
			clientId := "clientId-" + randomString(10)
			results := sendRequestsOneClient(clientId, *url,
				10, time.Duration(0.1*float32(time.Second)), 20*time.Second)

			fmt.Printf("Client %s results: %v\n", clientId, results)
		})
	}

	wg.Wait()
}
