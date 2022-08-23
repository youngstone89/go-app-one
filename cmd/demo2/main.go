package main

import (
	"fmt"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		return
	}
	fmt.Println("hostname: ", hostname)
}
