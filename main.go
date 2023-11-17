package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

var URL string = "http://localhost:80"
var NUM_REQ int
var NUM_CLIENTS int

var outchan chan float64 = make(chan float64)

func worker1(worker_id int) {
	client := http.Client{}
	var duration time.Duration
	for i := 1; i <= NUM_REQ; i++ {
		req, _ := http.NewRequest(http.MethodPut, URL, nil)
		req.Header.Add("Key", fmt.Sprintf("Key-%d-%d", worker_id, i))
		req.Header.Add("Val", fmt.Sprintf("Val-%d-%d", worker_id, i))

		start := time.Now()
		_, err := client.Do(req)
		duration += time.Since(start)

		if err != nil {
			fmt.Println("Error sending request.", err)
		}

		req, _ = http.NewRequest(http.MethodGet, URL, nil)
		req.Header.Add("Key", fmt.Sprintf("Key-%d-%d", worker_id, i))

		start = time.Now()
		_, err = client.Do(req)
		duration += time.Since(start)
		if err != nil {
			fmt.Println("Error sending request.", err)
		}

		req, _ = http.NewRequest(http.MethodDelete, URL, nil)
		req.Header.Add("Key", fmt.Sprintf("Key-%d-%d", worker_id, i))

		start = time.Now()
		_, err = client.Do(req)
		duration += time.Since(start)
		if err != nil {
			fmt.Println("Error sending request.", err)
		}
	}

	outchan <- float64(duration.Milliseconds())
}

func main() {

	NUM_REQ, _ = strconv.Atoi(os.Args[1])
	NUM_CLIENTS, _ = strconv.Atoi(os.Args[2])
	latency := 0.0

	start := time.Now()
	for id := 1; id <= NUM_CLIENTS; id++ {
		go worker1(id)
	}

	for id := 1; id <= NUM_CLIENTS; id++ {
		latency += <-outchan
	}

	duration := time.Since(start)
	close(outchan)
	TOTAL_REQ := NUM_CLIENTS * NUM_REQ * 3
	fmt.Printf("Total Requests: %d\n", TOTAL_REQ)
	fmt.Printf("Total throughput: %f req/s\n", float64(TOTAL_REQ)/duration.Seconds())
	fmt.Printf("Total latency: %f ms/req\n", latency/float64(NUM_CLIENTS))
}
