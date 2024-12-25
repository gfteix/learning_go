package concurrency

type WebsiteChecker func(string) bool

/*
Because the only way to start a goroutine is to put go in front of a function call, we often use anonymous functions when we want to start a goroutine.

---

Go can help us to spot race conditions with its built in race detector. To enable this feature, run the tests with the race flag: go test -race.

---

We can solve race conditions by coordinating our goroutines using channels.
Channels are a Go data structure that can both receive and send values.
These operations, along with their details, allow communication between different processes.

---


By sending the results into a channel, we can control the timing of each write into the results map, ensuring that it happens one at a time.
Although each of the calls of wc, and each send to the result channel, is happening concurrently inside its own process,
each of the results is being dealt with one at a time as we take values out of the result channel with the receive expression.
*/

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)} // send statement
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
