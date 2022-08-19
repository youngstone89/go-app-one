package main

import (
	"github.com/youngstone89/go-module-one/api/kafka"
)

func main() {
	kafkaClient := kafka.New()
	kafkaClient.Init()
}
