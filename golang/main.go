package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"gopkg.in/robfig/cron.v2"
)

func hello(name string) {
	message := fmt.Sprintf("Hi, %v", name)
	fmt.Println(message)
}

func runCronJobs1() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Seconds().Do(func() {
		hello("John Doe")
	})
	s.StartBlocking()
}

func runCronJobs2() {
	s := cron.New()
	s.AddFunc("@every 1s", func() {
		hello("John Doe")
	})
	s.Start()
}

func main() {
	fmt.Println("Using CronV2")
	runCronJobs2()
	fmt.Scanln()
}
