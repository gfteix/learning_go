package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

/*
We know we want our Countdown function to write data somewhere and io.Writer is the de-facto way of capturing that as an interface in Go.
In main we will send to os.Stdout so our users see the countdown printed to the terminal.
In test we will send to bytes.Buffer so our tests can capture what data is being generated.
*/

type Sleeper interface {
	Sleep()
}

const finalWord = "Go!"
const countdownStart = 3

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}
