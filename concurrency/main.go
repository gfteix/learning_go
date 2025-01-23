package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()

	responsech := make(chan string, 3)

	wg := &sync.WaitGroup{}

	// synchronously, the three functions together takes  290ms
	// using concurrency we can execute the process with the total time of the slowest function, that is 120ms
	// Note: this applies if the functions are independent and do not block each other.

	wg.Add(1)
	go fetchUserData(responsech, wg)
	wg.Add(1)
	go fetchUserRecommendation(responsech, wg)
	wg.Add(1)
	go fetchUserLikes(responsech, wg)

	wg.Wait()

	close(responsech)

	for resp := range responsech {
		fmt.Println(resp)
	}

	fmt.Println(time.Since(now))

}

func fetchUserData(responsech chan string, wg *sync.WaitGroup) {
	defer wg.Done() // Ensureing Done() is always called, even if an error occurs
	time.Sleep(80 * time.Millisecond)

	responsech <- "user data"
}

func fetchUserRecommendation(responsech chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(120 * time.Millisecond)

	responsech <- "user recommendation"
}

func fetchUserLikes(responsech chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(90 * time.Millisecond)

	responsech <- "user likes"
}
