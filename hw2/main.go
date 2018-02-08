package main

import (
	"log"
	"time"
)

// type job func(in, out chan interface{})

func ExecutePipeline(jobs ...job) {

	inChannel := make(chan interface{}, 3)

	for _, currentJob := range jobs {
		outChannel := make(chan interface{}, 3)
		go currentJob(inChannel, outChannel)
		inChannel = outChannel
	}

	time.Sleep(3 * time.Second)
}

func itemsJob(in, out chan interface{}) {
	items := []int{0, 1, 3, 4, 5}
	for _, item := range items {
		log.Println("item <-", item)
		out <- item
	}
}

func main() {
	jobs := []job{
		job(itemsJob),
	}
	ExecutePipeline(jobs...)
}
