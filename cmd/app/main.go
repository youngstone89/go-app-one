package main

import (
	"github.com/youngstone89/go-module-one/pkg/kafka"
)

func main() {
	kafkaClient := kafka.New()
	kafkaClient.Init()
}
