package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, fnCanel := context.WithCancel(context.Background())
	go func() {
		<-time.After(5 * time.Second)
		fmt.Println("caneling the context...")
		fnCanel()
	}()
out:
	for {
		select {
		case <-ctx.Done():
			break out
		default:
			fmt.Println("sleeping for 1 second...")
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("exiting the programm...")

}
