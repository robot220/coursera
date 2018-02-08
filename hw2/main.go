package main

import (
	"log"
	"time"
)

// type job func(in, out chan interface{})

func ExecutePipeline(jobs ...job) {

	inChannel := make(chan interface{}, 3)

	for _, value := range jobs {
		outChannel := make(chan interface{}, 3)
		go value(inChannel, outChannel)
		inChannel = outChannel
	}

	time.Sleep(3 * time.Second)
}

func main() {
	jobs := []job{
		job(func(in, out chan interface{}) {
			items := []int{0, 1, 3, 4, 5}
			for _, item := range items {
				log.Println("item <-", item)
				out <- item
			}
		}),
	}
	ExecutePipeline(jobs...)
}
