package _select

import (
	"fmt"
	"net/http"
	"time"
)

/*
## Synchronising processes

Why are we testing the speeds of the websites one after another when Go is great at concurrency? We should be able to check both at the same time.

We don't really care about the exact response times of the requests, we just want to know which one comes back first.

To do this, we're going to introduce a new construct called select which helps us synchronise processes really easily and clearly.
*/
var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		// including time.After in one of ours cases to prevent our system blocking forever.
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

/*
You'll recall from the concurrency chapter that you can wait for values to be sent to a channel with myVar := <-ch. This is a blocking call, as you're waiting for a value.

select allows you to wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.
*/

func ping(url string) chan struct{} {
	// we don't care about the result that is why we are using `chan struct{}` -> is the smallest data type available from a memory perspective
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
