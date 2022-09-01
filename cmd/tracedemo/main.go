package main

import (
	"os"
	"runtime/trace"
)

func main() {
	// start or stop writing the trace data to the standard error output
	trace.Start(os.Stderr)
	defer trace.Stop()
	// initialize unbuffered channel
	ch := make(chan int)

	// a goroutine that sends a number to the channel
	go func() {
		// send 42 to channel
		ch <- 1
	}()
	// read from channel
	<-ch
}
