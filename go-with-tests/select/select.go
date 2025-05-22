package _select

import (
	"net/http"
)

/*
## Synchronising processes

Why are we testing the speeds of the websites one after another when Go is great at concurrency? We should be able to check both at the same time.

We don't really care about the exact response times of the requests, we just want to know which one comes back first.

To do this, we're going to introduce a new construct called select which helps us synchronise processes really easily and clearly.
*/

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

/*
You'll recall from the concurrency chapter that you can wait for values to be sent to a channel with myVar := <-ch. This is a blocking call, as you're waiting for a value.

select allows you to wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.
*/

func ping(url string) chan struct{} {
	ch := make(chan struct{}) // chan struct{} is the smallest data type available from a memory perspective

	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
